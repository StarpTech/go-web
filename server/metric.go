package server

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Metric struct{}

type metricReport struct {
	Health string `json:"health"`
}

func (ctrl Metric) GetMetric(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := metricReport{"OK"}
		return c.JSON(http.StatusOK, m)
	}
}
