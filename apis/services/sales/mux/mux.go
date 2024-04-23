// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/ardanlabs/service/apis/services/sales/route/sys/checkapi"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
