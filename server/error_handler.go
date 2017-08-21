package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type InternalError struct {
	Code    string
	Message string
}

type HTTPResponseError struct {
	Code    int
	Message interface{}
	Stack   string
}

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		c.JSON(code, HTTPResponseError{Code: code, Message: he.Message, Stack: he.Error()})
	} else {
		c.JSON(code, InternalError{Code: "ErrInternal", Message: err.Error()})
	}
}
