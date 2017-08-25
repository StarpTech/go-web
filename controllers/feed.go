package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/store"
)

type Feed struct {
	Store  store.User
	Cache  store.Cache
	Config *config.Configuration
}

func (ctrl Feed) GetFeed(c echo.Context) error {
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return c.JSON(http.StatusOK, user)
}
