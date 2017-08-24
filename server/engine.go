package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v "gopkg.in/go-playground/validator.v9"
)

type Engine struct {
	Echo       *echo.Echo     // HTTP middleware
	config     *Configuration // Central configuration
	Logger     *log.Logger    // Global logger also for request logging
	greyLogger *gelf.TCPWriter
	db         *gorm.DB      // database connection
	cache      *redis.Client // redis cache connection
}

// NewEngine will create a new instance of the application
func NewEngine(config *Configuration) *Engine {
	engine := &Engine{}
	engine.Echo = echo.New()
	engine.config = config
	engine.cache = NewCache(config.RedisAddr, config.RedisPwd)
	engine.db = NewDB(config.Dialect, config.ConnectionString)

	// define validator
	engine.Echo.Validator = &Validator{validator: v.New()}

	//engine.Echo.Use(middleware.Recover())       // panic errors are thrown
	engine.Echo.Use(middleware.Logger())        // request logger
	engine.Echo.Use(middleware.BodyLimit("5M")) // limit body payload to 5MB
	engine.Echo.Use(middleware.Secure())        // provide protection against injection attacks
	engine.Echo.Use(middleware.RequestID())     // generate unique requestId

	// add custom error formatter
	engine.Echo.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := NewTemplateRenderer(engine.config.TemplateDir+"/layouts/*.html", engine.config.TemplateDir+"/*.html")
	engine.Echo.Renderer = renderer

	// add controllers
	userCtrl := new(User)
	feedCtrl := new(Feed)
	healthCtrl := new(Healthcheck)
	importCtrl := new(Importer)

	// api rest endpoints
	g := engine.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON(engine))

	// pages
	u := engine.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser(engine))
	u.GET("/:id/details", userCtrl.GetUserDetails(engine))

	// special endpoints
	engine.Echo.POST("/import", importCtrl.ImportUser(engine))
	engine.Echo.GET("/feed", feedCtrl.GetFeed(engine))

	// metric / health endpoint according to RFC 5785
	engine.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck(engine))
	engine.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	return engine
}

func (e *Engine) GetDB() *gorm.DB {
	return e.db
}

func (e *Engine) SetGrayLogger(g *gelf.TCPWriter) {
	e.greyLogger = g
}

// Start the http server
func (e *Engine) Start(addr string) error {
	return e.Echo.Start(addr)
}

// SetLogger set the logger instance for http server and internal
func (e *Engine) SetLogger(l *log.Logger) {
	e.Logger = l
	e.Echo.Logger = l
}

// ServeStaticFiles serve static files
func (e *Engine) ServeStaticFiles() {
	e.Echo.Static("/", e.config.AssetsBuildDir)
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (e *Engine) GracefulShutdown() {
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// close cache
	if e.cache != nil {
		cErr := e.cache.Close()
		if cErr != nil {
			e.Logger.Fatal(cErr)
		}
	}

	// close database connection
	if e.db != nil {
		dErr := e.db.Close()
		if dErr != nil {
			e.Logger.Fatal(dErr)
		}
	}

	// close greylogger tcp connection
	if e.greyLogger != nil {
		gErr := e.greyLogger.Close()
		if gErr != nil {
			e.Logger.Fatal(gErr)
		}
	}

	// shutdown http server
	if err := e.Echo.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
