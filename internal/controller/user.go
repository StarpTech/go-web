package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/starptech/go-web/internal/context"
	"github.com/starptech/go-web/internal/core/errors"
	"github.com/starptech/go-web/internal/models"
)

type (
	User          struct{}
	UserViewModel struct {
		Name string
		ID   string
	}
)

func (ctrl User) GetUser(c echo.Context) error {
	cc := c.(*context.AppContext)
	userID := c.Param("id")

	user := models.User{ID: userID}

	err := cc.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name: user.Name,
		ID:   user.ID,
	}

	return c.Render(http.StatusOK, "user.html", vm)

}

func (ctrl User) GetUserJSON(c echo.Context) error {
	cc := c.(*context.AppContext)
	userID := c.Param("id")

	user := models.User{ID: userID}

	err := cc.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name: user.Name,
		ID:   user.ID,
	}

	return c.JSON(http.StatusOK, vm)
}
