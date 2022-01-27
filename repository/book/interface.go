package book

import (
	_entities "sirclo/graphql/entities"
)

type Book interface {
	Get() ([]_entities.Book, error)
	Create(_entities.Book) (_entities.Book, error)
	GraphGet() ([]_entities.BookUserFormat, error)
	GraphGetByID(id int) (_entities.BookUserFormat, error)
}
