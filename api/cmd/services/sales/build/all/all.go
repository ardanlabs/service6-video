// Package all binds all the routes into the specified app.
package all

import (
	"github.com/ardanlabs/service/api/http/api/mux"
	"github.com/ardanlabs/service/api/http/domain/checkapi"
	"github.com/ardanlabs/service/api/http/domain/testapi"
	"github.com/ardanlabs/service/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	testapi.Routes(app, testapi.Config{
		Log:        cfg.Log,
		AuthClient: cfg.AuthClient,
	})
}
