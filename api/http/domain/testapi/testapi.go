package testapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/app/api/errs"
	"github.com/ardanlabs/service/foundation/web"
)

type api struct{}

func newAPI() *api {
	return &api{}
}

func (api *api) testError(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this message is trused")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func (api *api) testPanic(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		panic("WE ARE PANICKING!!!")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func (api *api) testAuth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
