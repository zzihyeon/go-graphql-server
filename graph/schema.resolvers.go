package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zzihyeon/go-graphql-server/graph/generated"
	"github.com/zzihyeon/go-graphql-server/graph/model"
	"github.com/zzihyeon/go-graphql-server/service/stocks"
)

func (r *mutationResolver) CreateStock(ctx context.Context, input model.NewStock) (*model.StandardResponse, error) {
	return stocks.Create(input), nil
}

func (r *mutationResolver) DeleteStock(ctx context.Context, input string) (*model.StandardResponse, error) {
	return stocks.Delete(input), nil
}

func (r *queryResolver) Stocks(ctx context.Context) (*model.StandardResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
