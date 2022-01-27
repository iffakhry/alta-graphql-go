package graph

import (
	"context"
	_bookRepo "sirclo/graphql/repository/book"
	_userRepo "sirclo/graphql/repository/user"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// db      *gorm.DB
	userRepo _userRepo.User
	bookRepo _bookRepo.Book
	// tmpList  []*_graphModel.User
	// Observer map[string]chan []*model.Person
	// Observer map[string]chan *_graphModel.User
	// mu sync.Mutex
}

func NewResolver(ur _userRepo.User, br _bookRepo.Book) *Resolver {
	return &Resolver{
		userRepo: ur,
		bookRepo: br,
		// tmpList:  []*_graphModel.User{},
		// Observer: map[string]chan *_graphModel.User{},
		// mu:       sync.Mutex{},
	}
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
