package userapi

import (
	"github.com/ardanlabs/service/api/http/api/mid"
	"github.com/ardanlabs/service/app/api/auth"
	"github.com/ardanlabs/service/app/api/authclient"
	"github.com/ardanlabs/service/app/domain/userapp"
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)
	ruleAuthorizeUser := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOrSubject)
	ruleAuthorizeAdmin := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOnly)

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandleFunc("GET /users", api.query, authen, ruleAdmin)
	app.HandleFunc("GET /users/{user_id}", api.queryByID, authen, ruleAuthorizeUser)
	app.HandleFunc("POST /users", api.create, authen, ruleAdmin)
	app.HandleFunc("PUT /users/role/{user_id}", api.updateRole, authen, ruleAuthorizeAdmin)
	app.HandleFunc("PUT /users/{user_id}", api.update, authen, ruleAuthorizeUser)
	app.HandleFunc("DELETE /users/{user_id}", api.delete, authen, ruleAuthorizeUser)
}
