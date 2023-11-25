package projectAPI

import (
	"net/http"
	"strconv"
	"worklogger/controllers/projects"

	"github.com/gin-gonic/gin"
)

func HandleListAll(c *gin.Context) {
	data, err := projects.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "N"})
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
	data, err := projects.Get(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

func HandlePost(c *gin.Context) {
	var project projects.Project

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := projects.Create(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New project created"})
}

func HandlePatch(c *gin.Context) {
	var project projects.Project

	queryId := c.Param("id")

	id, err := strconv.Atoi(queryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := projects.Update(id, project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
}

func HandleDelete(c *gin.Context) {
	queryId := c.Param("id")

	id, err := strconv.Atoi(queryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := projects.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
