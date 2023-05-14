package models

import (
	"backend/initializer"
	"log"
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

func (u *User) GetUser() (usr User, err error) {
	user := User{}
	row := initializer.Db.QueryRow("SELECT password , email FROM users WHERE email=$1", u.Email)
	if row.Err() != nil {
		return User{}, row.Err()
	}

	err = row.Scan(&user.Password, &user.Email)
	log.Println(err)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
