// Define all the routes here
package main

import "github.com/gin-gonic/gin"

func jsonMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
}

func initRoutes() {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.Use(jsonMiddleware)
		todoRoutes.Use(setUserMiddleware())
		todoRoutes.Use(logUserReqMiddleware())
		todoRoutes.GET("/", findTodos)
		todoRoutes.POST("/", createTodo)
		todoRoutes.PATCH("/:id", updateTodo)
		todoRoutes.DELETE("/:id", deleteTodo)
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.Use(jsonMiddleware)
		authRoutes.POST("/logins", handleAuth)
	}

}
