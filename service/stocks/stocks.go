package stocks

import (
	"log"

	"go-graphql-server/graph/model"
	"go-graphql-server/infrastructure/db/mongoDB"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(data model.NewStock) *model.StandardResponse {
	log.Println("!!create Stocks!!")
	return mongoDB.Create(data)
}

func Delete(name string) *model.StandardResponse {
	filter := bson.D{{Key: "name", Value: name}}
	return mongoDB.Delete(filter)
}
