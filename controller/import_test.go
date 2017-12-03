package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestImportUser(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/import", strings.NewReader(`{ "name": "bernd" }`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
