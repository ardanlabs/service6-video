// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/ardanlabs/service/apis/services/api/mid"
	"github.com/ardanlabs/service/apis/services/sales/route/sys/checkapi"
	"github.com/ardanlabs/service/business/api/auth"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(mux, auth)

	return mux
}
