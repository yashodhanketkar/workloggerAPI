package AuthAPI

import (
	"net/http"
	"time"
	"worklogger/controllers/users"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = "mysec"

func GenerateToken(name, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func HandleLogin(c *gin.Context) {
	var user users.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := users.Login(user); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateToken(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func HandleRegister(c *gin.Context) {
	var user users.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	if err := users.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register"})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"user":    user.Name,
	})
}
