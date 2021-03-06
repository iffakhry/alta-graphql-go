package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Books    []Book
}

type UpdateUserData struct {
	Name     *string
	Email    *string
	Password *string
}
