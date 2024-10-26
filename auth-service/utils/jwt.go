package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateJWT generates a JWT token for a given username with a 72-hour expiration.
func GenerateJWT(username string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(secret))
}
