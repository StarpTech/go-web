package context

import (
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/i18n"
	"github.com/starptech/go-web/store"
)

// Context is the new context in the request / response cycle
// We can use the db store, cache and central configuration
type AppContext struct {
	UserStore store.User
	Cache     store.Cache
	Config    *config.Configuration
	Loc       i18n.I18ner
}
