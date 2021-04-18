package mongoDB

import (
	"context"
	"log"

	"go-graphql-server/graph/model"
	"go-graphql-server/infrastructure/db/errors"

	"go.mongodb.org/mongo-driver/bson"
)

func deferFunc(code string, returnData *model.StandardResponse) {
	if r := recover(); r != nil {
		*returnData = errors.GetErrorResponse(code, nil)
	}
}

func Create(data model.NewStock) *model.StandardResponse {
	var returnData model.StandardResponse
	collection := Client.Database("test").Collection("stock")
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		returnData = errors.GetErrorResponse("Stock_Create_Failed", nil)
	} else {
		returnData = errors.GetErrorResponse("", insertResult)
	}
	defer deferFunc("Stock_Create_Panic", &returnData)
	return &returnData
}

func Read() {
	// collection := Client.Database("test").Collection("stock")
}

func Reads() {
	// collection := Client.Database("test").Collection("stock")
}

func Update() {
	// collection := Client.Database("test").Collection("stock")
}

func Delete(filter bson.D) *model.StandardResponse {
	log.Println("!!create Stocks Domain!!")
	var returnData model.StandardResponse
	collection := Client.Database("test").Collection("stock")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		returnData = errors.GetErrorResponse("Stock_Delete_Failed", nil)
	} else {
		returnData = errors.GetErrorResponse("", deleteResult)
	}
	defer deferFunc("Stock_Delete_Panic", &returnData)
	return &returnData
}
