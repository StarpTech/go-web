package context

import (
	"github.com/labstack/echo/v4"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/internal/i18n"
	"github.com/starptech/go-web/internal/store"
)

// AppContext is the new context in the request / response cycle
// We can use the db store, cache and central configuration
type AppContext struct {
	echo.Context
	UserStore store.User
	Cache     store.Cache
	Config    *config.Configuration
	Loc       i18n.I18ner
}
