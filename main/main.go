// Application entry point to bootstrap the webserver and db
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"godo/todoservice"
)

var router *gin.Engine

func main() {
	// Init db
	fmt.Println(os.Getenv("DB_CONN"), os.Getenv("DB_CLIENT"))
	db, err := gorm.Open(os.Getenv("DB_CLIENT"), os.Getenv("DB_CONN"))
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()
	fmt.Println("DB connected")

	// Init services
	todoservice.Init(db)

	// Init web server
	router = gin.Default()
	initRoutes()
	router.Run("0.0.0.0:8080")
}
