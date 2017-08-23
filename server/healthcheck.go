package server

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Healthcheck struct{}

type healthcheckReport struct {
	Health string `json:"health"`
}

func (ctrl Healthcheck) GetHealthcheck(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := healthcheckReport{"OK"}

		if db.DB().Ping() != nil {
			m.Health = "NOT"
			return c.JSON(http.StatusBadGateway, m)
		}

		return c.JSON(http.StatusOK, m)
	}
}
