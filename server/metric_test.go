package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestMetric(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/.well-known/metrics", nil)
	rec := httptest.NewRecorder()
	e.engine.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
