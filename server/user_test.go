package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestUnitUserCtrl(t *testing.T) {
	e := NewEngine()

	req := httptest.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.Echo.NewContext(req, rec)

	e.Echo.DefaultHTTPErrorHandler(errors.New("error"), c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestUserPage(t *testing.T) {
	e := NewEngine()
	req := httptest.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	e.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserDetailPage(t *testing.T) {
	e := NewEngine()
	req := httptest.NewRequest(echo.GET, "/users/1/details", nil)
	rec := httptest.NewRecorder()
	e.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
