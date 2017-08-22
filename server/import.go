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

	if errV := c.Validate(u).(validator.ValidationErrors); errV != nil {
		c.Logger().Error(errV.Error())
		return errors.New(errV.Error())
	}

	model := models.User{Name: u.Name}
	db := db.GetDB()

	if errM := db.Create(&model).Error; errM != nil {
		c.Logger().Errorf("entity could not be created %q", errM)
		return c.JSON(http.StatusBadRequest, u) // @TODO consistent error handling
	}

	return c.JSON(http.StatusOK, u)
}
