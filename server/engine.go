package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/starptech/go-web/config"
	v "gopkg.in/go-playground/validator.v9"
)

func NewEngine() *echo.Echo {
	c := config.GetConfig()
	e := echo.New()
	e.Debug = !c.IsProduction

	// define validator
	e.Validator = &Validator{validator: v.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

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
	e.Static("/", "public")

	return e
}
