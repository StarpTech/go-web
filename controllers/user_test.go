package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/models"
	"github.com/stretchr/testify/assert"
)

type UserFakeStore struct{}

func (s *UserFakeStore) First(m *models.User) error {
	return nil
}
func (s *UserFakeStore) Create(m *models.User) error {
	return nil
}
func (s *UserFakeStore) Ping() error {
	return nil
}

func TestUserPage(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUnitGetUserJson(t *testing.T) {

	s := echo.New()
	g := s.Group("/api")

	req := httptest.NewRequest(echo.GET, "/api/users/1", nil)
	rec := httptest.NewRecorder()

	userCtrl := &User{
		Config: e.config,
		Store:  &UserFakeStore{},
	}

	g.GET("/users/:id", userCtrl.GetUserJSON)
	s.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserDetailPage(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/users/1/details", nil)
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
