package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"sirclo/graphql/delivery/middlewares"
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
	} else {
		if input.Book != nil {
			_, err := r.bookRepo.Create(_entities.Book{
				Title:     input.Book.Title,
				Publisher: input.Book.Publisher,
				UserId:    responseData.ID,
			})
			if err != nil {
				return nil, errors.New("failed create book")
			}
		}
	}

	userResponseData := _graphModel.User{
		ID:    int(responseData.ID),
		Name:  responseData.Name,
		Email: responseData.Email,
	}
	return &userResponseData, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, set _graphModel.ChangeUser) (*_graphModel.SuccessResponse, error) {
	userData := _entities.UpdateUserData{
		Name:     set.Name,
		Email:    set.Email,
		Password: set.Password,
	}

	rowAffect, err := r.userRepo.Update(id, userData)
	if err != nil {
		return nil, err
	}
	var responseData _graphModel.SuccessResponse
	if rowAffect > 0 && err == nil {
		responseData.Message = "Update data success"
	} else {
		responseData.Message = "nothing changed"
	}
	return &responseData, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*_graphModel.SuccessResponse, error) {
	rowAffect, err := r.userRepo.Delete(id)
	if err != nil {
		return nil, err
	}
	var responseData _graphModel.SuccessResponse
	if rowAffect > 0 && err == nil {
		responseData.Message = "Delete data success"
	} else {
		responseData.Message = "nothing changed. Please make sure ID is correct"
	}
	return &responseData, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input _graphModel.NewBook) (*_graphModel.Book, error) {
	res, err := r.bookRepo.Create(_entities.Book{Title: input.Title, Publisher: input.Publisher, UserId: uint(input.Userid)})

	if err != nil {
		return nil, errors.New("not found")
	}
	convID := int(res.ID)

	return &_graphModel.Book{ID: convID, Title: res.Title, Publisher: res.Publisher}, nil
}

func (r *mutationResolver) BuatBook(ctx context.Context, title string, publisher string) (*_graphModel.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, set *_graphModel.ChangeBook) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context) ([]*_graphModel.Book, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convData := ctx.Value("EchoContextKey").(*middlewares.User)
		fmt.Println("id user", convData.Id)
	}

	responseData, err := r.bookRepo.GraphGet()
	fmt.Println(responseData)
	if err != nil {
		return nil, errors.New("not found")
	}
	booksData := []*_graphModel.Book{}
	for _, value := range responseData {
		bookResponseData := _graphModel.Book{
			ID:        value.ID,
			Title:     value.Title,
			Publisher: value.Publisher,
			Userid:    value.UserID,
			User: &_graphModel.User{
				ID:    value.UserID,
				Name:  value.Name,
				Email: value.Email,
			},
		}
		booksData = append(booksData, &bookResponseData)

	}
	// bookResponseData.ID = responseData.ID
	// bookResponseData.Title = responseData.Title
	// bookResponseData.Publisher = responseData.Publisher
	// usr := _graphModel.User{}
	// usr.ID = responseData.UserID
	// usr.Name = responseData.Name
	// usr.Email = responseData.Email
	// bookResponseData.User = &usr
	fmt.Println(booksData)
	return booksData, nil
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
		passwordData := v.Password
		userResponseData = append(userResponseData, &_graphModel.User{ID: convertID, Name: v.Name, Email: v.Email, Password: &passwordData})
	}

	return userResponseData, nil
}

func (r *queryResolver) UserByID(ctx context.Context, id int) (*_graphModel.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*_graphModel.LoginResponse, error) {
	id, name, token, err := r.authRepo.Login(email, password)
	if err != nil {
		return nil, err
	}
	return &_graphModel.LoginResponse{
		Message: "Login Success",
		ID:      &id,
		Name:    &name,
		Token:   &token,
	}, nil
}

func (r *queryResolver) BooksBySearch(ctx context.Context, data *_graphModel.BookData) (*_graphModel.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) BooksSearch(ctx context.Context, data *_graphModel.BookData) (*_graphModel.Book, error) {
	panic(fmt.Errorf("not implemented"))
}
