package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !isAuthenticated(c) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// LoggerMiddleware 日志记录中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求信息
		log.Println("Request:", c.Request.Method, c.Request.URL.Path)
		// 继续处理请求
		c.Next()
		// 记录响应信息
		log.Println("Response:", c.Writer.Status(), c.Writer.Size())
	}
}

// isAuthenticated 检查用户是否已认证
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
