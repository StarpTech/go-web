package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/db"
	"github.com/starptech/go-web/models"
)

type User struct{}

func (ctrl User) GetUser(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		c.Logger().Error("invalid user id")
		return c.String(http.StatusBadRequest, "")
	}

	db := db.GetDB()
	user := models.User{ID: uint64(userID)}

	if db.First(&user).Error != nil {
		c.Logger().Errorf("user with id %v could not be found", userID)
		return c.String(http.StatusNotFound, "")
	}

	return c.Render(http.StatusOK, "user.html", user)
}

func (ctrl User) GetUserDetails(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		c.Logger().Error("invalid user id")
		return c.String(http.StatusBadRequest, "") // @TODO consistent error handling
	}

	db := db.GetDB()
	user := models.User{ID: uint64(userID)}

	if db.First(&user).Error != nil {
		c.Logger().Errorf("user with id %v could not be found", userID)
		return c.String(http.StatusNotFound, "") // @TODO consistent error handling
	}

	return c.Render(http.StatusOK, "details.html", user)
}

func (ctrl User) GetUserJSON(c echo.Context) error {
	up := c.Param("id")
	userID, err := strconv.Atoi(up)

	if err != nil {
		c.Logger().Error("invalid user id")
		return c.String(http.StatusBadRequest, "") // @TODO consistent error handling
	}

	db := db.GetDB()
	user := models.User{ID: uint64(userID)}

	if db.First(&user).Error != nil {
		c.Logger().Errorf("user with id %v could not be found", userID)
		return c.String(http.StatusNotFound, "") // @TODO consistent error handling
	}

	return c.JSON(http.StatusOK, user)
}
