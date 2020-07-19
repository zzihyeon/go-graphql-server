package stocks

import (
	"log"

	"github.com/zzihyeon/go-graphql-server/domain/db/stock"
	"github.com/zzihyeon/go-graphql-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(data model.NewStock) *model.StandardResponse {
	log.Println("!!create Stocks!!")
	return stock.Create(data)
}

func Delete(name string) *model.StandardResponse {
	filter := bson.D{{"name", name}}
	return stock.Delete(filter)
}
