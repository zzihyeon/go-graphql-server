package types

// var StockType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Stock",
// 	Fields: graphql.Fields{
// 		"name":    &graphql.Field{Type: graphql.String},
// 		"persent": &graphql.Field{Type: graphql.String},
// 		"volume":  &graphql.Field{Type: graphql.Int},
// 		// "editions": &graphql.Field{Type: graphql.NewList(graphql.String)},
// 	},
// })

// var FavoriteType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Favorite",
// 	Fields: graphql.Fields{
// 		"name": &graphql.Field{Type: graphql.String},
// 		"date": &graphql.Field{Type: graphql.String},
// 		// "editions": &graphql.Field{Type: graphql.NewList(graphql.String)},
// 	},
// })

// var ReasonType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Reason",
// 	Fields: graphql.Fields{
// 		"name":   &graphql.Field{Type: graphql.String},
// 		"date":   &graphql.Field{Type: graphql.String},
// 		"stocks": &graphql.Field{Type: graphql.NewList(StockType)},
// 	},
// })

type StockType struct {
	Name   string `json:"name"`
	Per    string `json:"name"`
	Volume string `json:"name"`
}
