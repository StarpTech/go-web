package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/models"
)

type (
	User          struct{}
	UserViewModel struct {
		Name      string
		PublicDir string
	}
)

func (ctrl User) GetUser(e *Engine) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.config.AssetsPublicDir,
		}

		return c.Render(http.StatusOK, "user.html", vm)
	}
}

func (ctrl User) GetUserDetails(e *Engine) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(b)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.config.AssetsPublicDir,
		}

		return c.Render(http.StatusOK, "details.html", vm)
	}
}

func (ctrl User) GetUserJSON(e *Engine) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = e.GetDB().First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		vm := UserViewModel{
			Name:      user.Name,
			PublicDir: e.config.AssetsPublicDir,
		}

		return c.JSON(http.StatusOK, vm)
	}
}
