// This tool will output a demo JWT token
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User map[string]interface{} `json:"user"`
	jwt.StandardClaims
}

func generateJWT() (string, error) {
	claims := Claims{
		User: map[string]interface{}{
			"id":   "local:1",
			"name": "John Doe",
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(), // Token expires in 1 year
			Subject:   "local:1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func main() {
	token, err := generateJWT()
	if err != nil {
		fmt.Println("Failed to generate token:", err)
		return
	}
	fmt.Println("Generated JWT Token:")
	fmt.Println(token)
}
