package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/models"
	"github.com/starptech/go-web/server"
)

type (
	User          struct{}
	UserViewModel struct {
		Name      string
		PublicDir string
	}
)

func (ctrl User) GetUser(e *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := server.Boom{Code: server.InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := server.Boom{Code: server.UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.Config.AssetsPublicDir,
		}

		return c.Render(http.StatusOK, "user.html", vm)
	}
}

func (ctrl User) GetUserDetails(e *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := server.Boom{Code: server.InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(b)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := server.Boom{Code: server.UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.Config.AssetsPublicDir,
		}

		return c.Render(http.StatusOK, "details.html", vm)
	}
}

func (ctrl User) GetUserJSON(e *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := server.Boom{Code: server.InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := server.Boom{Code: server.UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.Config.AssetsPublicDir,
		}

		return c.JSON(http.StatusOK, vm)
	}
}
