package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	switch v := err.(type) {
	case *echo.HTTPError:
		e := c.JSON(v.Code, v)
		if e != nil {
			c.Logger().Error("error handler: json encoding", e)
		}
	default:
		e := c.JSON(code, boom{Code: InternalError, Message: "Bad implementation"})
		if e != nil {
			c.Logger().Error("error handler: json encoding", e)
		}
	}
}
