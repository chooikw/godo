package main

import (
	"bytes"

	"fmt"

	"io"

	"github.com/gin-gonic/gin"
)

func logUserReqMiddleware() gin.HandlerFunc {
	// Logs user's requests
	return func(c *gin.Context) {

		var bodyString string
		user, _ := c.Get("user")

		if c.Request.Method == "PATCH" || c.Request.Method == "POST" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(500, gin.H{"error": "Error reading request body"})
				c.Abort()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the request body for downstream handlers
			bodyString = fmt.Sprintf(" Req Body: %v", string(body))
		}

		fmt.Printf("[Req] User #%v%v\n", user.(UserClaims).Id, bodyString)

		c.Next()
	}
}
