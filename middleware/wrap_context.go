package middleware

import (
	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
)

func WrapContext(cc *context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return next(cc)
		}
	}
}
