package rest

import (
	"api/adapters/rest/routes"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	r := routes.CreateRoutes()
	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
