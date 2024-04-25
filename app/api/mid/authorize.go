package mid

import (
	"context"
	"errors"

	"github.com/ardanlabs/service/app/api/errs"
	"github.com/ardanlabs/service/business/api/auth"
)

// ErrInvalidID represents a condition where the id is not a uuid.
var ErrInvalidID = errors.New("ID is not in its proper form")

// Authorize executes the specified role and does not extract any domain data.
func Authorize(ctx context.Context, auth *auth.Auth, rule string, handler Handler) error {
	userID, err := GetUserID(ctx)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	if err := auth.Authorize(ctx, GetClaims(ctx), userID, rule); err != nil {
		return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, claims[%v] rule[%v]: %s", GetClaims(ctx).Roles, rule, err)
	}

	return handler(ctx)
}
