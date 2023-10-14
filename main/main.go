package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Init db

	// Init web server
	router = gin.Default()
	initRoutes()
	router.Run("0.0.0.0:8080")
}
