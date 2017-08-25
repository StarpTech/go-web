package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/core/errors"
	"github.com/starptech/go-web/models"
	"github.com/starptech/go-web/store"
)

type (
	User struct {
		Store  store.User
		Cache  store.Cache
		Config *config.Configuration
	}
	UserViewModel struct {
		Name      string
		PublicDir string
	}
)

func (ctrl User) GetUser(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, "invalid user id", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Store.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Config.AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "user.html", vm)

}

func (ctrl User) GetUserDetails(c echo.Context) error {

	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, "invalid user id", err)
		c.Logger().Error(b)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Store.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Config.AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "details.html", vm)

}

func (ctrl User) GetUserJSON(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := errors.NewBoom(errors.InvalidUserID, "invalid user id", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Store.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Config.AssetsPublicDir,
	}

	return c.JSON(http.StatusOK, vm)
}
