package auth

import (
	_middlewares "sirclo/graphql/delivery/middlewares"
	_entities "sirclo/graphql/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) Login(email string, password string) (id int, name string, token string, err error) {
	// var err error
	// var authToken string
	user := _entities.User{}
	tx := a.db.Where("email=? AND password=?", email, password).First(&user)
	if tx.Error != nil {
		return id, name, token, tx.Error
	}

	token, err = _middlewares.CreateToken(int(user.ID))
	if err != nil {
		return int(user.ID), user.Name, "", err
	}
	// if er := config.DB.Save(user).Error; er != nil {
	// 	return nil, err
	// }
	return int(user.ID), user.Name, token, nil

	// if username == "admin" && password == "admin" {
	// 	claims := jwt.MapClaims{}
	// 	claims["authorized"] = true
	// 	claims["id"] = 1
	// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// 	claims["name"] = username

	// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 	return token.SignedString([]byte("R4HASIA"))
	// }
	// return "", errors.New("failed login!")
}
