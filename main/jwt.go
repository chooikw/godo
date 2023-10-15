// This is a helper to function to generate JWT token
package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	User User `json:"user"`
	jwt.StandardClaims
}

func generateJWTForUser(user User) (string, error) {
	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(), // Token expires in 1 year
			Subject:   user.Id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
