package server

import (
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	v "gopkg.in/go-playground/validator.v9"
)

func NewEngine() *echo.Echo {
	c := config.GetConfig()
	e := echo.New()

	logger := log.New("server")

	if c.GrayLogAddr != "" {
		gelfWriter, err := gelf.NewUDPWriter(c.GrayLogAddr)
		if err != nil {
			logger.Fatalf("gelf.NewWriter: %s", err)
		}
		// Log to greylog and stderr
		logger.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
	}

	// enable colors for beautying the beast
	if c.IsProduction {
		logger.EnableColor()
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.ERROR)
	}

	e.Logger = logger

	// define validator
	e.Validator = &Validator{validator: v.New()}
	e.Use(middleware.Recover())
	// add custom error formatter
	e.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := NewTemplateRenderer(c.TemplateDir+"/layouts/*.html", c.TemplateDir+"/*.html")
	e.Renderer = renderer

	// add controllers
	userCtrl := new(User)
	feedCtrl := new(Feed)
	importCtrl := new(Importer)

	// add api endpoints
	g := e.Group("/api")
	g.GET("/:id", userCtrl.GetUserJSON)

	// add routes
	e.GET("/:id", userCtrl.GetUser)
	e.GET("/:id/details", userCtrl.GetUserDetails)
	e.POST("/import", importCtrl.ImportUser)

	// add feed
	e.GET("/feed", feedCtrl.GetFeed)

	// add static files
	e.Static("/", c.PublicDir)

	return e
}
