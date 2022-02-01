package user

import (
	"fmt"
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

func (p *UserRepository) Create(user _entities.User) (_entities.User, error) {
	if err := p.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (p *UserRepository) Update(id int, user _entities.UpdateUserData) (int, error) {
	dataUpdate := make(map[string]interface{})
	if user.Name != nil {
		dataUpdate["name"] = user.Name
	}
	if user.Email != nil {
		dataUpdate["email"] = user.Email
	}
	if user.Password != nil {
		dataUpdate["password"] = user.Password
	}
	tx := p.db.Model(&_entities.User{}).Where("id = ?", id).Updates(dataUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	fmt.Println("row", tx.RowsAffected)
	return int(tx.RowsAffected), nil
}

func (p *UserRepository) Delete(id int) (int, error) {
	tx := p.db.Delete(&_entities.User{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
