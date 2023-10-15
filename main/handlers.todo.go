// TODO routes handlers
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func findTodos(c *gin.Context) {
	// TODO
}

func createTodo(c *gin.Context) {
	// TODO
	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	// TODO
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	// TODO
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}
