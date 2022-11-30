package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/EronAlves1996/GoTeste/graph/generated"
	"github.com/EronAlves1996/GoTeste/graph/model"
)

// Verify is the resolver for the verify field.
func (r *queryResolver) Verify(ctx context.Context, password string, rules []*model.Rule) (*model.Return, error) {
	panic(fmt.Errorf("not implemented: Verify - verify"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
