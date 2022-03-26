package routes

import (
	"api/adapters/rest/routes/endpoints"
	"api/models/User"

	"github.com/gorilla/mux"
)

func CreateRoutes(userRepository User.Repository) *mux.Router {
	r := mux.NewRouter()
	return endpoints.ConfigRoutes(r, userRepository)
}
