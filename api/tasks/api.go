package taskAPI

import (
	"fmt"
	"net/http"
	"strconv"
	"worklogger/controllers/tasks"

	"github.com/gin-gonic/gin"
)

func HandleListAll(c *gin.Context) {
	data, err := tasks.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Interal error"})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No data in table"})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

func HandleGet(c *gin.Context) {
	queryId := c.Param("id")

	id, err := strconv.Atoi(queryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	data, err := tasks.Get(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

var task tasks.Task

func HandlePost(c *gin.Context) {

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := tasks.Create(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New task created"})
}

func HandlePatch(c *gin.Context) {
	var task tasks.Task

	queryId := c.Param("id")

	id, err := strconv.Atoi(queryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := tasks.Update(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func HandleDelete(c *gin.Context) {
	queryId := c.Param("id")

	id, err := strconv.Atoi(queryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tasks.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
