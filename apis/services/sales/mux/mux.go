// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"encoding/json"
	"net/http"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	}

	mux.HandleFunc("GET /test", h)

	return mux
}
