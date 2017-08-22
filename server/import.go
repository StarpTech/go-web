package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/db"
	"github.com/starptech/go-web/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type Importer struct{}

type UserEntity struct {
	Name string `json:"name" validate:"required"`
}

func (ctrl Importer) ImportUser(c echo.Context) error {
	u := new(UserEntity)

	if errB := c.Bind(u); errB != nil {
		c.Logger().Errorf("entity could not be bound %q", errB)
		return c.String(http.StatusBadRequest, "") // @TODO consistent error handling
	}

	errV := c.Validate(u)

	if errV != nil {
		err := errV.(validator.ValidationErrors)
		c.Logger().Error(err.Error())
		return errors.New(err.Error())
	}

	model := models.User{Name: u.Name}
	db := db.GetDB()

	if errM := db.Create(&model).Error; errM != nil {
		c.Logger().Errorf("entity could not be created %q", errM)
		return c.JSON(http.StatusBadRequest, u) // @TODO consistent error handling
	}

	return c.JSON(http.StatusOK, u)
}
