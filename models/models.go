package models

import (
	"backend/initializer"
	"fmt"
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
	fmt.Println("in db")
	row := initializer.Db.QueryRow(` INSERT INTO users
    (email , password , first_name,last_name)
	VALUES
	($1,$2,$3,$4)`,
		u.Email, u.Password, u.FirstName, u.LastName)

	if row.Err() != nil {
		return http.StatusInternalServerError, row.Err()
	}

	return http.StatusCreated, nil
}
