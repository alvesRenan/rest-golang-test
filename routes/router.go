package routes

import "github.com/gorilla/mux"

// NewRouter returns a mux.Router configured
// with all the application routes
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Func).Methods(r.Method)
	}

	return router
}
