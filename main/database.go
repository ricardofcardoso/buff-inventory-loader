package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var mongoClient *mongo.Client
var inventoryCollection *mongo.Collection

func initDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongodbUri := os.Getenv("MONGODB_URI")

	if mongodbUri != "" {
		opts := options.Client().ApplyURI(mongodbUri).SetServerAPIOptions(serverAPI)

		var err error
		mongoClient, err = mongo.Connect(context.Background(), opts)
		if err != nil {
			log.Fatal(err)
		}

		if err = mongoClient.Database("BUFF").RunCommand(context.Background(), bson.D{{"ping", 1}}).Err(); err != nil {
			log.Fatal(err)
		}

		inventoryCollection = mongoClient.Database("BUFF").Collection("inventory")

		fmt.Println("Successfully connected to MongoDB!")
	} else {
		log.Fatal("Please set the env variable MONGODB_URI with your MongoDB URI")
	}
}
