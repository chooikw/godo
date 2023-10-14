package main

func initRoutes() {
	//router.Use(setUserStatus())

	todoRoutes := router.Group("/todos")
	{
		todoRoutes.Use(setUserMiddleware())
		todoRoutes.Use(logUserReqMiddleware())
		todoRoutes.GET("/", findTodos)
		todoRoutes.POST("/", createTodo)
		todoRoutes.PATCH("/:id", updateTodo)
		todoRoutes.DELETE("/:id", deleteTodo)
	}
}
