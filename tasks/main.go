package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var col *mongo.Collection

func mongodbConnection() {
	uri := "mongodb://root:root@mongo:27017/"
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	col = client.Database("tastdb").Collection("taskdata")

}

func main() {
	mongodbConnection()

}
