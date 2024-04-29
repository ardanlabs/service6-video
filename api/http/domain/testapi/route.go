package testapi

import (
	"github.com/ardanlabs/service/api/http/api/mid"
	"github.com/ardanlabs/service/app/api/auth"
	"github.com/ardanlabs/service/app/api/authclient"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	athAdminOnly := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI()

	app.HandleFunc("GET /testerror", api.testError)
	app.HandleFunc("GET /testpanic", api.testPanic)
	app.HandleFunc("GET /testauth", api.testAuth, authen, athAdminOnly)
}
