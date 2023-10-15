// Middleware to decode the JWT token, and set user from the payload into gin Context
package main

import (
	"godo/authservice"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func setUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenString := authHeader[7:] // Remove the prefix Bearer

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user information from the token into the context
		claims, _ := token.Claims.(jwt.MapClaims)
		userId, _ := claims["user"].(map[string]interface{})["id"].(string)
		userName, _ := claims["user"].(map[string]interface{})["name"].(string)

		user := authservice.User{
			Id:   userId,
			Name: userName,
		}
		c.Set("user", user)

		c.Next()
	}
}
