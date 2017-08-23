package server

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/starptech/go-web/models"
)

type User struct{}

func (ctrl User) GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = db.First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		return c.Render(http.StatusOK, "user.html", user)
	}
}

func (ctrl User) GetUserDetails(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(b)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = db.First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		return c.Render(http.StatusOK, "details.html", user)
	}
}

func (ctrl User) GetUserJSON(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		up := c.Param("id")
		userID, err := strconv.Atoi(up)

		if err != nil {
			b := boom{Code: InvalidUserID, Message: "invalid user id", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, b)
		}

		user := models.User{ID: uint64(userID)}

		err = db.First(&user).Error

		if err != nil {
			b := boom{Code: UserNotFound, Message: "user could not be found", Details: err}
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, b)
		}

		return c.JSON(http.StatusOK, user)
	}
}
