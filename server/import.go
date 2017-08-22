package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type Importer struct{}

type UserEntity struct {
	Name string `json:"name"`
}

func (ctrl Importer) ImportUser(c echo.Context) error {
	u := new(UserEntity)
	err := c.Bind(u)

	if err != nil {
		c.Logger().Fatalf("entity could not be bound %q", err)
		return c.String(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, u)
}
