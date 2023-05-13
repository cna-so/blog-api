package helpers

import (
	"backend/models"
	"errors"
)

func ValidateUser(user models.User) (models.User, error) {
	if len(user.Password) == 0 {
		return models.User{}, errors.New("password is required")
	}
	if len(user.Password) <= 7 {
		return models.User{}, errors.New("password must be more than 8 length")
	}
	if len(user.Email) == 0 {
		return models.User{}, errors.New("email is required")
	}
	if len(user.Email) <= 5 {
		return models.User{}, errors.New("email must be more than 5 length")
	}
	return user, nil
}
