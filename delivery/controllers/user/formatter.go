package user

import (
	_entities "sirclo/graphql/entities"
)

type UserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponseFormat struct {
	Id    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func (req *UserRequestFormat) ToEntity() *_entities.User {
	return &_entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func FromEntity(entities _entities.User) UserResponseFormat {
	return UserResponseFormat{
		Id:    entities.ID,
		Name:  entities.Name,
		Email: entities.Email,
		// UpdatedAt: domain.UpdatedAt,
		// DeletedAt: domain.DeletedAt,
	}
}
