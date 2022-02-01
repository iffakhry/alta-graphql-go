package user

import (
	_entities "sirclo/graphql/entities"
)

type User interface {
	Get() ([]_entities.User, error)
	Create(_entities.User) (_entities.User, error)
	Update(id int, user _entities.UpdateUserData) (int, error)
	Delete(id int) (int, error)
}
