package routes

import (
	"api/adapters/rest/routes/endpoints"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	return endpoints.ConfigRoutes(r)
}
