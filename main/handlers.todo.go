// TODO routes handlers
package main

import (
	"net/http"

	"godo/todoservice"

	"github.com/gin-gonic/gin"

	"strconv"
)

type CreateTodoData struct {
	Data todoservice.Todo `json:"data"`
}

type UpdateTodoData struct {
	Data todoservice.UpdateInput `json:"data"`
}

func findTodos(c *gin.Context) {
	todos := todoservice.FindMany()
	c.IndentedJSON(http.StatusOK, gin.H{"data": todos})
}

func createTodo(c *gin.Context) {
	var createData CreateTodoData
	err := c.BindJSON(&createData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	todo, createErr := todoservice.Create(createData.Data)
	if createErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unable to create data"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": todo})
}

func updateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var updateData UpdateTodoData
	err := c.BindJSON(&updateData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	updatedTodo, updateErr := todoservice.Update(id, updateData.Data)
	if updateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": updatedTodo})
}

func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := todoservice.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}
