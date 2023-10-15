// Application entry point to bootstrap the webserver and db
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var router *gin.Engine

func main() {
	// Init db
	db, err := gorm.Open(os.Getenv("DB_CLIENT"), os.Getenv("DB_CONN"))
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
	defer db.Close()
	fmt.Println("DB connected")

	// Init web server
	router = gin.Default()
	initRoutes()
	router.Run("0.0.0.0:8080")
}
