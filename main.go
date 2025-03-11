package main

import (
	"log"
	AuthAPI "worklogger/api/auth"
	projectAPI "worklogger/api/projects"
	taskAPI "worklogger/api/tasks"
	"worklogger/db"
	auth "worklogger/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// db setup
	db.InitDB()

	// router setup
	router := gin.New()

	apiGroup := router.Group("/api/v1")

	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/login", AuthAPI.HandleLogin)
		authGroup.POST("/register", AuthAPI.HandleRegister)
	}

	apiGroup.Use(auth.AuthMiddleware())
	{
		taskGroup := apiGroup.Group("/tasks")
		{
			taskGroup.GET("/", taskAPI.HandleListAll)
			taskGroup.POST("/", taskAPI.HandlePost)
			taskGroup.GET("/:id", taskAPI.HandleGet)
			taskGroup.PATCH("/:id", taskAPI.HandlePatch)
			taskGroup.DELETE("/:id", taskAPI.HandleDelete)
		}
		projectGroup := apiGroup.Group("/projects")
		{
			projectGroup.GET("/", projectAPI.HandleListAll)
			projectGroup.POST("/", projectAPI.HandlePost)
			projectGroup.GET("/:id", projectAPI.HandleGet)
			projectGroup.PATCH("/:id", projectAPI.HandlePatch)
			projectGroup.DELETE("/:id", projectAPI.HandleDelete)
		}
	}
	log.Fatal("Server initilization failed: ", router.Run("127.0.0.1:8000"))
}
