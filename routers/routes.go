package routers

import (
	"github.com/blackaay/gin-start/internal/middleware"
	UserController "github.com/blackaay/gin-start/internal/module/user/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	auth := r.Group("/api/auth")
	auth.Use(middleware.AuthMiddleware(), middleware.LoggerMiddleware())
	{
		auth.POST("/user", UserController.Create)
		auth.GET("/user", UserController.Get)
		auth.POST("/logout", UserController.Logout)
	}
	auth.Use(middleware.LoggerMiddleware())
	{
		auth.POST("/login", UserController.Login)
	}
}
