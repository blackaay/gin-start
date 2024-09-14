package routers

import (
	"github.com/blackaay/gin-start/internal/controllers"
	"github.com/blackaay/gin-start/internal/middleware"
	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化路由
func InitRoutes(r *gin.Engine) {
	// 应用局部中间件
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/login", controllers.LoginUser)    // 登录用户
		auth.POST("/logout", controllers.LogoutUser) // 注销用户
	}
}
