package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/app/api/mid"
	"github.com/ardanlabs/service/business/api/auth"
	"github.com/ardanlabs/service/foundation/web"
)

// Authorization validates a JWT from the `Authorization` header.
func Authorization(auth *auth.Auth) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Authorization(ctx, auth, r.Header.Get("authorization"), hdl)
		}

		return h
	}

	return m
}
