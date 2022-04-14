package routes

import (
	"api/adapters/rest/routes/endpoints"
	"api/infra/repository/user"

	"github.com/gorilla/mux"
)

func CreateRoutes(userRepository user.Repository) *mux.Router {
	r := mux.NewRouter()
	return endpoints.ConfigRoutes(r, userRepository)
}
