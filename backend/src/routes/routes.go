package routes

import (
	"api/src/routes/endpoints"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	return endpoints.ConfigRoutes(r)
}
