package main

import (
	"github.com/blackaay/gin-start/internal/middleware"
	"github.com/blackaay/gin-start/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	routers.InitRoutes(router)
	port := "8080"

	// 启动 HTTP 服务器
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
