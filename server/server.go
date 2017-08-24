package server

import (
	"context"
	"os"
	"os/signal"
	"time"

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
	Echo   *echo.Echo            // HTTP middleware
	Config *config.Configuration // Configuration
	db     *models.Model         // Database connection
	cache  *redis.Client         // Redis cache connection
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
func (s *Server) GetDB() *models.Model {
	return s.db
}

// GetCache returns the current redis client
func (s *Server) GetCache() *redis.Client {
	return s.cache
}

// Start the http server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}

// ServeStaticFiles serve static files for development purpose
func (s *Server) ServeStaticFiles() {
	s.Echo.Static("/", s.Config.AssetsBuildDir)
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// close cache
	if s.cache != nil {
		cErr := s.cache.Close()
		if cErr != nil {
			s.Echo.Logger.Fatal(cErr)
		}
	}

	// close database connection
	if s.db != nil {
		dErr := s.db.Close()
		if dErr != nil {
			s.Echo.Logger.Fatal(dErr)
		}
	}

	// shutdown http server
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}
}
