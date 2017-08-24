package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestUserPage(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	e.engine.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserDetailPage(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/users/1/details", nil)
	rec := httptest.NewRecorder()
	e.engine.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
