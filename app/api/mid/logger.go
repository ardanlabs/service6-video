package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			return err
		}

		return h
	}

	return m
}
