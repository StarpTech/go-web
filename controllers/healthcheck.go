package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/store"
)

type Healthcheck struct {
	Store  store.User
	Cache  store.Cache
	Config *config.Configuration
}

type healthcheckReport struct {
	Health  string          `json:"health"`
	Details map[string]bool `json:"details"`
}

// GetHealthcheck return the current functional state of the application
func (ctrl Healthcheck) GetHealthcheck(c echo.Context) error {
	m := healthcheckReport{Health: "OK"}

	dbCheck := ctrl.Store.Ping()
	cacheCheck := ctrl.Cache.Ping()

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
