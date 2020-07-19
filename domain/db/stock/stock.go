package stock

import (
	"context"

	"github.com/zzihyeon/go-graphql-server/graph/model"
	"github.com/zzihyeon/go-graphql-server/infrastructure/db"
	"go.mongodb.org/mongo-driver/bson"
)

func deferFunc(code string, returnData *model.StandardResponse) {
	if r := recover(); r != nil {
		*returnData = GetErrorResponse(code, nil)
	}
}

func Create(data model.Stock) model.StandardResponse {
	var returnData model.StandardResponse
	client := db.Setup()
	collection := client.Database("test").Collection("stock")
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		returnData = GetErrorResponse("Stock_Create_Failed", nil)
	} else {
		returnData = GetErrorResponse("", insertResult)
	}
	defer deferFunc("Stock_Create_Panic", &returnData)
	return returnData
}

func Read() {
	client := db.Setup()
	collection := client.Database("test").Collection("stock")
}

func Reads() {
	client := db.Setup()
	collection := client.Database("test").Collection("stock")
}

func Update() {
	client := db.Setup()
	collection := client.Database("test").Collection("stock")
}

func Delete(filter bson.D) model.StandardResponse {
	var returnData model.StandardResponse
	client := db.Setup()
	collection := client.Database("test").Collection("stock")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		returnData = GetErrorResponse("Stock_Delete_Failed", nil)
	} else {
		returnData = GetErrorResponse("", insertResult)
	}
	defer deferFunc("Stock_Delete_Panic", &returnData)
	return returnData
}
