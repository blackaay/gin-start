package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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
		startTime := time.Now()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		latency := time.Since(startTime)
		logEntry := fmt.Sprintf("[GIN] %s | %d | %v | %s | %s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path)
		log.Println(logEntry)
		// 继续处理请求
		c.Next()
	}
}

func isAuthenticated(c *gin.Context) bool {
	// 实现具体的认证逻辑
	return true
}
