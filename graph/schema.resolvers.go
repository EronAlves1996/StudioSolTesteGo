package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/EronAlves1996/GoTeste/graph/generated"
	"github.com/EronAlves1996/GoTeste/graph/model"
	"github.com/EronAlves1996/GoTeste/services"
)

// Verify is the resolver for the verify field.
func (r *queryResolver) Verify(ctx context.Context, password string, rules []*model.Rule) (*model.Return, error) {
	verify := services.ProcessPassword(password, rules)
	return verify, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
