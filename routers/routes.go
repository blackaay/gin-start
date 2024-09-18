package routers

import (
	"github.com/blackaay/gin-start/internal/middleware"
	UserController "github.com/blackaay/gin-start/internal/module/user/controllers"
	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化路由
func InitRoutes(r *gin.Engine) {
	// 应用局部中间件
	auth := r.Group("/api/auth")
	auth.Use(middleware.AuthMiddleware(), middleware.LoggerMiddleware())
	{
		auth.GET("/login", UserController.LoginUser)    // 登录用户
		auth.POST("/logout", UserController.LogoutUser) // 注销用户
	}
}
