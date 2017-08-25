package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
)

type Healthcheck struct{}

type healthcheckReport struct {
	Health  string          `json:"health"`
	Details map[string]bool `json:"details"`
}

// GetHealthcheck return the current functional state of the application
func (ctrl Healthcheck) GetHealthcheck(c echo.Context) error {
	cc := c.(*context.Context)
	m := healthcheckReport{Health: "OK"}

	dbCheck := cc.UserStore.Ping()
	cacheCheck := cc.Cache.Ping()

	if dbCheck != nil {
		m.Health = "NOT"
		m.Details["db"] = false
	}

	if cacheCheck != nil {
		m.Health = "NOT"
		m.Details["cache"] = false
	}

	return cc.JSON(http.StatusOK, m)
}
