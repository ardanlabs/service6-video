package checkapi

import (
	"github.com/ardanlabs/service/apis/services/api/mid"
	"github.com/ardanlabs/service/business/api/auth"
	"github.com/ardanlabs/service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, a *auth.Auth) {
	authen := mid.Authorization(a)
	athAdminOnly := mid.Authorize(a, auth.RuleAdminOnly)

	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
	app.HandleFunc("GET /testauth", liveness, authen, athAdminOnly)
}
