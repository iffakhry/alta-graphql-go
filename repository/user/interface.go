package user

import (
	_entities "sirclo/graphql/entities"
)

type User interface {
	Get() ([]_entities.User, error)
	Create(_entities.User) (_entities.User, error)
}
