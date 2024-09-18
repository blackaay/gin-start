package user_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser 用户注册
func LoginUser(c *gin.Context) {
	// 处理注册逻辑
	c.JSON(http.StatusOK, gin.H{"message": "登陆成功"})
}

// GetUserByID 获取用户信息
func LogoutUser(c *gin.Context) {
	// 处理获取用户信息逻辑
	c.JSON(http.StatusOK, gin.H{"message": "退出成功"})
}
