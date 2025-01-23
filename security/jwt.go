package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, role string, id int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"name": username,
		"role": role,
		"exp":  time.Now().Add(2 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
