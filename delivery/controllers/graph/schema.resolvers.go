package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	_entities "sirclo/graphql/entities"
	_graphModel "sirclo/graphql/entities/graph/model"
	"sirclo/graphql/util/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input _graphModel.NewUser) (*_graphModel.User, error) {
	userData := _entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	responseData, err := r.userRepo.Create(userData)
	if err != nil {
		return nil, errors.New("failed create user")
	}

	userResponseData := _graphModel.User{
		ID:    int(responseData.ID),
		Name:  responseData.Name,
		Email: responseData.Email,
	}
	return &userResponseData, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*_graphModel.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookByID(ctx context.Context, id int) (*_graphModel.Book, error) {
	responseData, err := r.bookRepo.GraphGetByID(id)
	fmt.Println(responseData)
	if err != nil {
		return nil, errors.New("not found")
	}
	bookResponseData := _graphModel.Book{
		ID:        responseData.ID,
		Title:     responseData.Title,
		Publisher: responseData.Publisher,
		Userid:    responseData.UserID,
		User: &_graphModel.User{
			ID:    responseData.UserID,
			Name:  responseData.Name,
			Email: responseData.Email,
		},
	}
	// bookResponseData.ID = responseData.ID
	// bookResponseData.Title = responseData.Title
	// bookResponseData.Publisher = responseData.Publisher
	// usr := _graphModel.User{}
	// usr.ID = responseData.UserID
	// usr.Name = responseData.Name
	// usr.Email = responseData.Email
	// bookResponseData.User = &usr
	fmt.Println(bookResponseData)
	return &bookResponseData, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*_graphModel.User, error) {
	responseData, err := r.userRepo.Get()

	if err != nil {
		return nil, errors.New("not found")
	}

	userResponseData := []*_graphModel.User{}

	for _, v := range responseData {
		convertID := int(v.ID)
		userResponseData = append(userResponseData, &_graphModel.User{ID: convertID, Name: v.Name, Email: v.Email, Password: &v.Password})
	}

	return userResponseData, nil
}

func (r *queryResolver) UserByID(ctx context.Context) (*_graphModel.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
