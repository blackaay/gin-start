package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// auth 权限验证
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !isAuthenticated(c) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// 日志记录中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("logger-middleware")
		// 继续处理请求
		c.Next()
		log.Println("logger-middleware-after")
	}
}

func isAuthenticated(c *gin.Context) bool {
	// 实现具体的认证逻辑
	name := c.DefaultQuery("name", "")
	log.Println("Request:", name)
	if name != "" {
		return true
	} else {
		return false
	}
}
