package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongo() *mongo.Client {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))

	if err != nil {
		log.Println(err)
		panic(err.Error())
		return nil
	}

	return client
}
