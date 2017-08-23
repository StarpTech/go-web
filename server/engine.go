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
	"github.com/starptech/go-web/config"
	v "gopkg.in/go-playground/validator.v9"
)

type Engine struct {
	Echo   *echo.Echo
	Config *config.Configuration
	Logger *log.Logger
	Db     *gorm.DB
}

func NewEngine() *Engine {

	config := config.GetConfig()
	db := NewDB(config.Dialect, config.ConnectionString)
	logger := NewLogger(config.GrayLogAddr, config.IsProduction)

	engine := &Engine{}
	engine.Db = db
	engine.Config = config
	engine.Echo = echo.New()
	engine.Logger = logger
	engine.Echo.Logger = logger

	// define validator
	engine.Echo.Validator = &Validator{validator: v.New()}
	engine.Echo.Use(middleware.Recover())
	engine.Echo.Use(middleware.Logger())

	// add custom error formatter
	engine.Echo.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := NewTemplateRenderer(engine.Config.TemplateDir+"/layouts/*.html", engine.Config.TemplateDir+"/*.html")
	engine.Echo.Renderer = renderer

	// add controllers
	userCtrl := new(User)
	feedCtrl := new(Feed)
	importCtrl := new(Importer)

	g := engine.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON(db))

	u := engine.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser(db))
	u.GET("/:id/details", userCtrl.GetUserDetails(db))
	u.POST("/import", importCtrl.ImportUser(db))
	u.GET("/feed", feedCtrl.GetFeed(db))

	return engine
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
