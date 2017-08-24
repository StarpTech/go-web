package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/boom"
	"github.com/starptech/go-web/models"
)

type (
	User struct {
		Context CtrlContext
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
		b := boom.New(boom.InvalidUserID, "invalid user id", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Context.GetDB().First(&user).Error

	if err != nil {
		b := boom.New(boom.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Context.GetConfig().AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "user.html", vm)

}

func (ctrl User) GetUserDetails(c echo.Context) error {

	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := boom.New(boom.InvalidUserID, "invalid user id", err)
		c.Logger().Error(b)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Context.GetDB().First(&user).Error

	if err != nil {
		b := boom.New(boom.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Context.GetConfig().AssetsPublicDir,
	}

	return c.Render(http.StatusOK, "details.html", vm)

}

func (ctrl User) GetUserJSON(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		b := boom.New(boom.InvalidUserID, "invalid user id", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, b)
	}

	user := models.User{ID: uint64(userID)}

	err = ctrl.Context.GetDB().First(&user).Error

	if err != nil {
		b := boom.New(boom.UserNotFound, "user could not be found", err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name:      user.Name,
		PublicDir: ctrl.Context.GetConfig().AssetsPublicDir,
	}

	return c.JSON(http.StatusOK, vm)
}
