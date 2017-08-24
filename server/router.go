package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	v "gopkg.in/go-playground/validator.v9"
)

func NewRouter(server *Server) *echo.Echo {
	config := server.config
	e := echo.New()
	// define validator
	e.Validator = &Validator{validator: v.New()}

	e.Use(middleware.Recover())       // panic errors are thrown
	e.Use(middleware.Logger())        // request logger
	e.Use(middleware.BodyLimit("5M")) // limit body payload to 5MB
	e.Use(middleware.Secure())        // provide protection against injection attacks
	e.Use(middleware.RequestID())     // generate unique requestId

	// add custom error formating
	e.HTTPErrorHandler = HTTPErrorHandler

	// Add html templates with go template syntax
	renderer := newTemplateRenderer(config.LayoutDir, config.TemplateDir)
	e.Renderer = renderer

	return e
}
