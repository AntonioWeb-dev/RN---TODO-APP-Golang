package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

var ctx = context.TODO()

func Init(url, dbname string) {
	clientOptions := options.Client().ApplyURI("mongo://root:example@localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	userCollection = client.Database("todoapp").Collection("users")
}
