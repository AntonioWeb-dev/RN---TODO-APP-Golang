package endpoints

import (
	"api/adapters/rest/middlewares"
	"api/infra/repository/user"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	Controller     func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

func ConfigRoutes(r *mux.Router, userRepository user.Repository) *mux.Router {
	routes := InitUserRoutes(userRepository)

	for _, router := range routes {
		if router.Authentication {
			r.HandleFunc(
				router.URI,
				middlewares.Authorization(router.Controller),
			).Methods(router.Method)
		}
		r.HandleFunc(router.URI, router.Controller).Methods(router.Method)
	}
	return r
}
