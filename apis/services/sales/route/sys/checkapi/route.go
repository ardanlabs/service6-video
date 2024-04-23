package checkapi

import "net/http"

// Routes adds specific routes for this group.
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
}
