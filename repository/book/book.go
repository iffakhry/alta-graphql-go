package book

import (
	"fmt"

	_entities "sirclo/graphql/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Get() ([]_entities.Book, error) {
	var books []_entities.Book
	// var tmp []_entities.User
	qry := br.db.Raw("Select books.* from books join users on users.ID = books.user_id").Scan(&books)
	// qry := br.db.Raw(`Select books.*, users.* from books join users on users.ID = books.user_id`).Scan(&books)
	// br.db.Joins("Persons").Find(&books)

	if err := qry.Error; err != nil {
		return nil, err
	}

	fmt.Println(qry.Statement)
	fmt.Println(qry.Statement.Preloads["Books"]...)
	fmt.Println(books)
	return books, nil
}

func (br *BookRepository) Create(book _entities.Book) (_entities.Book, error) {
	if err := br.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (br *BookRepository) GraphGetByID(id int) (_entities.BookUserFormat, error) {
	// type res struct {
	// 	ID        int
	// 	Title     string
	// 	Publisher string
	// 	UserID    int
	// 	Name      string
	// 	Email     string
	// 	Password  string
	// }
	var tmp _entities.BookUserFormat
	qry := br.db.Raw(`Select books.ID, books.Title, books.publisher, users.ID as 'UserID', users.name, users.email, users.password from books join users on users.ID = books.user_id where books.ID = ?`, id).Scan(&tmp)

	if err := qry.Error; err != nil {
		return tmp, err
	}

	return tmp, nil
	// return &_graphModel.Book{ID: &tmp.ID,
	// 	Title:     tmp.Title,
	// 	Publisher: tmp.Publisher,
	// 	Userid: &_graphModel.User{
	// 		ID:       &tmp.UserID,
	// 		Name:     tmp.Name,
	// 		Email:    tmp.Email,
	// 		Password: tmp.Password}}, nil
}
