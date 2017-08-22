package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/db"
	"github.com/starptech/go-web/models"
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

	model := models.User{Name: u.Name}
	db := db.GetDB()
	if db.Create(&model).Error != nil {
		c.Logger().Fatalf("entity could not be created %q", err)
		return c.JSON(http.StatusBadRequest, u)
	}

	return c.JSON(http.StatusOK, u)
}
