package rest

import (
	"api/adapters/rest/routes"
	"api/infra/database"
	"api/models/User"
	"context"
	"fmt"
	"log"
	"net/http"
)

var ctx = context.TODO()

func Init() {
	db := database.Init("mongodb://root:example@localhost:27017/", "todoapp", ctx)
	userRepository := User.InitRepo(ctx, db.UserCollection)

	r := routes.CreateRoutes(userRepository)
	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
