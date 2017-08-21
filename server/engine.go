package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/starptech/go-web/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// add custom error formatter
	e.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := NewTemplateRenderer("templates/layouts/*.html", "templates/*.html")
	e.Renderer = renderer

	// add controllers
	userCtrl := new(controllers.User)
	feedCtrl := new(controllers.Feed)

	// add routes
	e.GET("/:id", userCtrl.GetUser)
	e.GET("/:id/details", userCtrl.GetUserDetails)

	// add api endpoints
	g := e.Group("/api")
	g.GET("/:id", userCtrl.GetUserJSON)

	// add feed
	e.GET("/feed", feedCtrl.GetFeed)

	// add static files
	e.Static("/", "assets")

	return e
}
