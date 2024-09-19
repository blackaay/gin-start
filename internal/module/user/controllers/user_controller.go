package user_controllers

import (
	"github.com/blackaay/gin-start/internal/module/user/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var req validator.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	if err := validator.ValidateUserCreate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get"})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}
