package context

import (
	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/store"
)

type Context struct {
	echo.Context
	UserStore store.User
	Cache     store.Cache
	Config    *config.Configuration
}
