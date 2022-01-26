package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	_graphModel "sirclo/graphql/entities/graph/model"
	"sirclo/graphql/util/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input _graphModel.NewUser) (*_graphModel.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*_graphModel.User, error) {
	responseData, err := r.userRepo.Get()

	if err != nil {
		return nil, errors.New("not found")
	}

	userResponseData := []*_graphModel.User{}

	for _, v := range responseData {
		convertID := int(v.ID)
		userResponseData = append(userResponseData, &_graphModel.User{ID: &convertID, Name: v.Name, Email: v.Email, Password: v.Password})
	}

	return userResponseData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
