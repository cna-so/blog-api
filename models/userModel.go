package models

import (
	"backend/initializer"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name,omitempty" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}
type UserLogin struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

func (u *User) CreateUser() (statusCode int, err error) {
	row := initializer.Db.QueryRow(` INSERT INTO users
    (email , password ,role, first_name,last_name , created_at , updated_at)
	VALUES
	($1,$2,$3,$4,$5,$6,$7)`,
		u.Email, u.Password, "user", u.FirstName, u.LastName, time.Now(), time.Now())

	if row.Err() != nil {
		return http.StatusInternalServerError, row.Err()
	}

	return http.StatusCreated, nil
}

func (u *UserLogin) GetUser() (usr User, err error) {
	user := User{}
	row := initializer.Db.QueryRow("SELECT password , email ,role FROM users WHERE email=$1", u.Email)
	if row.Err() != nil {
		return User{}, row.Err()
	}

	err = row.Scan(&user.Password, &user.Email, &user.Role)
	log.Println(err)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
