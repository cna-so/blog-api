package authorization

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJwtToken(email, role string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  email,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, err
}
