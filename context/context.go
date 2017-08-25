package context

import (
	"github.com/labstack/echo"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/store"
)

// Context is the new context in the request / reponse cycle
type Context struct {
	echo.Context
	UserStore store.User
	Cache     store.Cache
	Config    *config.Configuration
}
