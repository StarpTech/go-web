package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/server"
)

type Healthcheck struct{}

type healthcheckReport struct {
	Health  string          `json:"health"`
	Details map[string]bool `json:"details"`
}

// GetHealthcheck return the current functional state of the application
func (ctrl Healthcheck) GetHealthcheck(e *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := healthcheckReport{Health: "OK"}

		dbCheck := e.GetDB().DB.DB().Ping()
		cacheCheck := e.GetCache().Ping().Err()

		if dbCheck != nil {
			m.Health = "NOT"
			m.Details["db"] = false
		}

		if cacheCheck != nil {
			m.Health = "NOT"
			m.Details["cache"] = false
		}

		return c.JSON(http.StatusOK, m)
	}
}
