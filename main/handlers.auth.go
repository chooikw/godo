package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"godo/authservice"
)

type AuthReq struct {
	Data struct {
		Strategy         string `json:"strategy"`
		VerificationCode string `json:"verificationCode"`
	} `json:"data" binding:"required"`
}

func handleAuth(c *gin.Context) {
	var reqData AuthReq
	err := c.BindJSON(&reqData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	var res gin.H
	switch reqData.Data.Strategy {
	case "github":
		githubRes := authservice.GithubVerifierUrl()
		fmt.Println(githubRes)
		res = gin.H{"data": githubRes}
	case "githubVerify":
		user, err := authservice.GithubVerifyCode(reqData.Data.VerificationCode)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		token, _ := generateJWTTokenForUser(user)
		res = gin.H{"data": token}
	default:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unknown strategy"})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
