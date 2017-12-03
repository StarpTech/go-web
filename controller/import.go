package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
	"github.com/starptech/go-web/core/errors"
	"github.com/starptech/go-web/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type Importer struct{}

type UserEntity struct {
	Name string `json:"name" validate:"required"`
}

func (ctrl Importer) ImportUser(c echo.Context) error {
	cc := c.(*context.AppContext)
	u := new(UserEntity)

	if errB := c.Bind(u); errB != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), errB)
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

	if errM := cc.UserStore.Create(&model); errM != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), errM)
		c.Logger().Error(errM)
		return c.JSON(http.StatusBadRequest, b)
	}

	return c.JSON(http.StatusOK, u)

}
