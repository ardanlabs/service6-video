package tests

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	authbuild "github.com/ardanlabs/service/api/cmd/services/auth/build/all"
	salesbuild "github.com/ardanlabs/service/api/cmd/services/sales/build/all"
	"github.com/ardanlabs/service/api/http/api/apitest"
	"github.com/ardanlabs/service/api/http/api/mux"
	"github.com/ardanlabs/service/app/api/auth"
	"github.com/ardanlabs/service/app/api/authclient"
	"github.com/ardanlabs/service/business/api/dbtest"
	"github.com/ardanlabs/service/foundation/docker"
)

var c *docker.Container

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	var err error

	c, err = dbtest.StartDB()
	if err != nil {
		return 1, err
	}
	defer dbtest.StopDB(c)

	return m.Run(), nil
}

func startTest(t *testing.T, testName string) *apitest.Test {
	db := dbtest.NewDatabase(t, c, testName)

	// -------------------------------------------------------------------------

	auth, err := auth.New(auth.Config{
		Log:       db.Log,
		KeyLookup: &apitest.KeyStore{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// -------------------------------------------------------------------------

	server := httptest.NewServer(mux.WebAPI(mux.Config{
		Log:  db.Log,
		Auth: auth,
		DB:   db.DB,
	}, authbuild.Routes()))

	logFunc := func(ctx context.Context, msg string, v ...any) {
		db.Log.Info(ctx, msg, v...)
	}

	authClient := authclient.New(server.URL, logFunc)

	// -------------------------------------------------------------------------

	mux := mux.WebAPI(mux.Config{
		Log:        db.Log,
		AuthClient: authClient,
		DB:         db.DB,
	}, salesbuild.Routes())

	return apitest.New(db, auth, mux)
}
