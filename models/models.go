package models

import (
	"backend/initializer"
	"errors"
	"net/http"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"`
}

func (u *User) CreateUser() (statusCode int, err error) {
	if len(u.Email) < 5 || len(u.Password) < 8 {
		return http.StatusNoContent, errors.New("the (email , password) are require")
	}

	row := initializer.Db.QueryRow("INSERT INTO users (email , password , first_name,last_name) VALUES ($1,$2,$3,$4)", u.Email, u.Password, u.FirstName, u.LastName)

	if row.Err() != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
