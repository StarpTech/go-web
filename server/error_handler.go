package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	switch v := err.(type) {
	case *echo.HTTPError:
		c.JSON(v.Code, v)
	default:
		c.JSON(code, v)
	}
}
