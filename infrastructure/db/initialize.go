package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go-graphql-server/infrastructure/db/mongoDB"
	"go-graphql-server/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize() {
	var config types.EnvConfig
	data, err := ioutil.ReadFile("config/development.json")
	json.Unmarshal(data, &config)
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Mongo)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	mongoDB.Client = client
	fmt.Println("Connected to MongoDB!")
}