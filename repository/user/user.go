package user

import (
	_entities "sirclo/graphql/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

//query untuk get all users
func (p *UserRepository) Get() ([]_entities.User, error) {
	var users []_entities.User
	if err := p.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
