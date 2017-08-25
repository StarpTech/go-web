package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
)

type Feed struct{}

func (ctrl Feed) GetFeed(c echo.Context) error {
	cc := c.(*context.Context)
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return cc.JSON(http.StatusOK, user)
}
