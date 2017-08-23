package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v "gopkg.in/go-playground/validator.v9"
)

type Engine struct {
	Echo   *echo.Echo     // HTTP middleware
	Config *Configuration // Central configuration
	Logger *log.Logger    // Global logger also for request logging
	Db     *gorm.DB       // database connection
}

// NewEngine will create a new instance of the application
func NewEngine(config *Configuration) *Engine {
	engine := &Engine{}
	engine.Echo = echo.New()
	engine.Config = config
	engine.Db = NewDB(config.Dialect, config.ConnectionString)

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
	renderer := NewTemplateRenderer(engine.Config.TemplateDir+"/layouts/*.html", engine.Config.TemplateDir+"/*.html")
	engine.Echo.Renderer = renderer

	// add controllers
	userCtrl := new(User)
	feedCtrl := new(Feed)
	healthCtrl := new(Healthcheck)
	importCtrl := new(Importer)

	g := engine.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON(engine.Db, engine.Config))

	u := engine.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser(engine.Db, engine.Config))
	u.GET("/:id/details", userCtrl.GetUserDetails(engine.Db, engine.Config))
	u.POST("/import", importCtrl.ImportUser(engine.Db, engine.Config))
	u.GET("/feed", feedCtrl.GetFeed(engine.Db, engine.Config))

	// metric / health endpoint according to RFC 5785
	engine.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck(engine.Db))
	engine.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	return engine
}

// Start the http server
func (e *Engine) Start(addr string) error {
	return e.Echo.Start(addr)
}

// SetLogger set the logger instance for http server and internal
func (e *Engine) SetLogger(logger *log.Logger) {
	e.Logger = logger
	e.Echo.Logger = logger
}

// ServeStaticFiles serve static files
func (e *Engine) ServeStaticFiles() {
	e.Echo.Static("/", e.Config.AssetsBuildDir)
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (e *Engine) GracefulShutdown() {
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Echo.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
