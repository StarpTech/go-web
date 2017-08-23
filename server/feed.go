package server

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Feed struct{}

func (ctrl Feed) GetFeed(db *gorm.DB, config *Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user struct {
			Name string `json:"name"`
		}
		user.Name = "Peter"
		return c.JSON(http.StatusOK, user)
	}
}
