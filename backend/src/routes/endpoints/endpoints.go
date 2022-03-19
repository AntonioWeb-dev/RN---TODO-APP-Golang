package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	Controller     func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := userEndpoints

	for _, router := range routes {
		r.HandleFunc(router.URI, router.Controller).Methods(router.Method)
	}
	return r
}
