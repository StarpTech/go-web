package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/server"
)

type Feed struct{}

func (ctrl Feed) GetFeed(e *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user struct {
			Name string `json:"name"`
		}
		user.Name = "Peter"
		return c.JSON(http.StatusOK, user)
	}
}
