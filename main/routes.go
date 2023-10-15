// Define all the routes here
package main

import "github.com/gin-gonic/gin"

func initRoutes() {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.Use(func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			c.Next()
		})
		todoRoutes.Use(setUserMiddleware())
		todoRoutes.Use(logUserReqMiddleware())
		todoRoutes.GET("/", findTodos)
		todoRoutes.POST("/", createTodo)
		todoRoutes.PATCH("/:id", updateTodo)
		todoRoutes.DELETE("/:id", deleteTodo)
	}
}
