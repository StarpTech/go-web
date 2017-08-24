package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Healthcheck struct {
	Context CtrlContext
}

type healthcheckReport struct {
	Health  string          `json:"health"`
	Details map[string]bool `json:"details"`
}

// GetHealthcheck return the current functional state of the application
func (ctrl Healthcheck) GetHealthcheck(c echo.Context) error {
	m := healthcheckReport{Health: "OK"}

	dbCheck := ctrl.Context.GetDB().DB.DB().Ping()
	cacheCheck := ctrl.Context.GetCache().Ping().Err()

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
