// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Userid    int    `json:"userid"`
	User      *User  `json:"user"`
}

type BookData struct {
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

type LoginResponse struct {
	Message string  `json:"message"`
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Token   *string `json:"token"`
}

type NewBook struct {
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Userid    *int   `json:"userid"`
}

type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Book     *NewBook `json:"book"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}
