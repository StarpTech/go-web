package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct{}

func (ctrl User) GetUser(c echo.Context) error {
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return c.Render(http.StatusOK, "user.html", user)
}

func (ctrl User) GetUserDetails(c echo.Context) error {
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return c.Render(http.StatusOK, "details.html", user)
}

func (ctrl User) GetUserJSON(c echo.Context) error {
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return c.JSON(http.StatusOK, user)
}
