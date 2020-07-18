package graphql

// import (
// 	"go-graphql-server/types"

// 	"github.com/graphql-go/graphql"
// )

// var schemaConfig = graphql.SchemaConfig{
// 	Mutation: rootMutation,
// }

// var rootMutation = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "RootMutation",
// 	Fields: graphql.Fields{
// 		"getStocks": &graphql.Field{
// 			Type:        types.StockType, // the return type for this field
// 			Description: "get Stocks by name",
// 			Args: graphql.FieldConfigArgument{
// 				"name": &graphql.ArgumentConfig{
// 					Type: graphql.NewList(graphql.String),
// 				},
// 			},
// 			// Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 			// 	log.Printf("Args: %v", params.Args)
// 			// 	jsonString, err := json.Marshal(params.Args)
// 			// 	if err != nil {
// 			// 		fmt.Println("Error encoding JSON")
// 			// 		return nil, nil
// 			// 	}

// 			// 	author := Author{}
// 			// 	json.Unmarshal([]byte(jsonString), &author)

// 			// 	return author, nil
// 			// },
// 		},
// 	},
// })
