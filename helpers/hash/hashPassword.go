package HashPassword

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword create a hash password from raw password
func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

// ComparePassWithHash  check the password is right or not
func ComparePassWithHash(hashPass, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	if err != nil {
		return true
	}
	return false
}
