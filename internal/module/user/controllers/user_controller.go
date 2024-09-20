package user_controllers

import (
	"github.com/blackaay/gin-start/internal/module/common"
	"github.com/blackaay/gin-start/internal/module/user/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var data validator.User
	if err := validator.ValidateUserCreate(c, &data); err == false {
		return
	}
	common.ResponseSuccess(c, "User created successfully")
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
