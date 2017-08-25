package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/core/errors"
	"github.com/starptech/go-web/models"
	"github.com/starptech/go-web/store"
	validator "gopkg.in/go-playground/validator.v9"
)

type Importer struct {
	Store  store.User
	Cache  store.Cache
	Config *config.Configuration
}

type UserEntity struct {
	Name string `json:"name" validate:"required"`
}

func (ctrl Importer) ImportUser(c echo.Context) error {

	u := new(UserEntity)

	if errB := c.Bind(u); errB != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, "invalid user model", errB)
		c.Logger().Error(errB)
		return c.JSON(http.StatusBadRequest, b)
	}

	errV := c.Validate(u)

	if errV != nil {
		err := errV.(validator.ValidationErrors)
		c.Logger().Error(err.Error())
		return err
	}

	model := models.User{Name: u.Name}

	if errM := ctrl.Store.Create(&model); errM != nil {
		b := errors.NewBoom(errors.EntityCreationError, "user could not be created", errM)
		c.Logger().Error(errM)
		return c.JSON(http.StatusBadRequest, b)
	}

	return c.JSON(http.StatusOK, u)

}
