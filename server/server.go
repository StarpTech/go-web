package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/cache"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/models"
	v "gopkg.in/go-playground/validator.v9"
)

type Server struct {
	Echo       *echo.Echo            // HTTP middleware
	Config     *config.Configuration // Configuration
	Logger     *log.Logger           // logger also used for request logging
	greyLogger *gelf.TCPWriter       // Graylogger logger
	db         *models.Model         // Database connection
	cache      *redis.Client         // Redis cache connection
}

// NewServer will create a new instance of the application
func NewServer(config *config.Configuration) *Server {
	server := &Server{}
	server.Echo = echo.New()
	server.Config = config
	server.db = models.NewModel()
	server.cache = cache.NewCache(config)

	err := server.db.OpenWithConfig(config)
	if err != nil {
		log.Errorf("gorm: could not connect to db %q", err)
	}

	// define validator
	server.Echo.Validator = &Validator{validator: v.New()}

	server.Echo.Use(middleware.Recover())       // panic errors are thrown
	server.Echo.Use(middleware.Logger())        // request logger
	server.Echo.Use(middleware.BodyLimit("5M")) // limit body payload to 5MB
	server.Echo.Use(middleware.Secure())        // provide protection against injection attacks
	server.Echo.Use(middleware.RequestID())     // generate unique requestId

	// add custom error formating
	server.Echo.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := newTemplateRenderer(server.Config.TemplateDir+"/layouts/*.html", server.Config.TemplateDir+"/*.html")
	server.Echo.Renderer = renderer

	return server
}

// GetDB returns gorm (ORM)
func (e *Server) GetDB() *models.Model {
	return e.db
}

// GetCache returns the current redis client
func (e *Server) GetCache() *redis.Client {
	return e.cache
}

// SetGrayLogger set the graylogger
func (e *Server) SetGrayLogger(g *gelf.TCPWriter) {
	e.greyLogger = g
}

// Start the http server
func (e *Server) Start(addr string) error {
	return e.Echo.Start(addr)
}

// SetLogger set the logger instance for http server and internal
func (e *Server) SetLogger(l *log.Logger) {
	e.Logger = l
	e.Echo.Logger = l
}

// ServeStaticFiles serve static files for development purpose
func (e *Server) ServeStaticFiles() {
	e.Echo.Static("/", e.Config.AssetsBuildDir)
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (e *Server) GracefulShutdown() {
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
