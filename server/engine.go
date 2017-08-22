package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/starptech/go-web/config"
)

func NewEngine() *echo.Echo {
	c := config.GetConfig()
	e := echo.New()
	e.Debug = !c.IsProduction
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// add custom error formatter
	e.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := NewTemplateRenderer("../templates/layouts/*.html", "../templates/*.html")
	e.Renderer = renderer

	// add controllers
	userCtrl := new(User)
	feedCtrl := new(Feed)
	importCtrl := new(Importer)

	// add routes
	e.GET("/:id", userCtrl.GetUser)
	e.GET("/:id/details", userCtrl.GetUserDetails)
	e.POST("/import", importCtrl.ImportUser)

	// add api endpoints
	g := e.Group("/api")
	g.GET("/:id", userCtrl.GetUserJSON)

	// add feed
	e.GET("/feed", feedCtrl.GetFeed)

	// add static files
	e.Static("/", "assets")

	return e
}
