package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
	"github.com/starptech/go-web/core/errors"
	"github.com/starptech/go-web/models"
)

type (
	User          struct{}
	UserViewModel struct {
		Name      string
		PublicDir string
	}
)

func (ctrl User) GetUser(c echo.Context) error {
	app := c.Get("app").(*context.AppContext)
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, errors.ErrorText(errors.InvalidUserID), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = app.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: app.Config.AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "user.html", vm)

}

func (ctrl User) GetUserDetails(c echo.Context) error {
	app := c.Get("app").(*context.AppContext)
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, errors.ErrorText(errors.InvalidUserID), err)
		c.Logger().Error(b)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = app.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: app.Config.AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "details.html", vm)

}

func (ctrl User) GetUserJSON(c echo.Context) error {
	app := c.Get("app").(*context.AppContext)
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, errors.ErrorText(errors.InvalidUserID), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = app.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: app.Config.AssetsPublicDir,
	}

	return c.JSON(http.StatusOK, vm)
}
