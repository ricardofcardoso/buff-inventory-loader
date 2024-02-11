package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var mongoClient *mongo.Client
var inventoryCollection *mongo.Collection

func initDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongodbUri := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(mongodbUri).SetServerAPIOptions(serverAPI)

	var err error
	mongoClient, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	if err = mongoClient.Database("BUFF").RunCommand(context.Background(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	inventoryCollection = mongoClient.Database("BUFF").Collection("inventory")

	fmt.Println("Successfully connected to MongoDB!")
}
