package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string
	Publisher string
	UserId    uint
}

type BookUserFormat struct {
	ID        int
	Title     string
	Publisher string
	UserID    int
	Name      string
	Email     string
	Password  string
}
