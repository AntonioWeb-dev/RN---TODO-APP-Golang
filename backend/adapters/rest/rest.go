package rest

import (
	"api/adapters/rest/routes"
	"api/infra/database"
	"api/infra/repository/user"
	"context"
	"fmt"
	"log"
	"net/http"
)

var ctx = context.TODO()

func Init() {
	db := database.Init("mongodb://root:example@localhost:27018/", "todoapp", ctx)
	userRepository := user.InitRepo(ctx, db.UserCollection)

	r := routes.CreateRoutes(userRepository)
	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
