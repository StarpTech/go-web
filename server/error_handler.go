package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/starptech/go-web/boom"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	switch v := err.(type) {
	case *echo.HTTPError:
		err := c.JSON(v.Code, v)
		if err != nil {
			c.Logger().Error("error handler: json encoding", err)
		}
	default:
		e := boom.New(boom.InternalError, "Bad implementation", nil)
		err := c.JSON(code, e)
		if err != nil {
			c.Logger().Error("error handler: json encoding", err)
		}
	}
}
