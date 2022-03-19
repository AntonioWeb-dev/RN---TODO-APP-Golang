package main

import (
	"api/src/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.CreateRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
