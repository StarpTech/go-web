package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/.well-known/health-check", nil)
	rec := httptest.NewRecorder()
	e.engine.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
