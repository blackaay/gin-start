package main

import (
	"github.com/blackaay/gin-start/configs"
	"github.com/blackaay/gin-start/database/migrations"
	"github.com/blackaay/gin-start/internal/middleware"
	"github.com/blackaay/gin-start/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	logFile *os.File
)

func main() {
	// 日志设置
	f, err := createDailyLogFile(false)
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(f, os.Stdout))
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 启动日志文件轮换的定时任务
	go checkAndRotateLogFile()

	//执行 migrate
	go migrations.Migrate()

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	routers.InitRoutes(router)
	port := configs.GetSettingString("HTTP_PORT")
	// 启动 HTTP 服务器
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func createDailyLogFile(nextDay bool) (*os.File, error) {
	var logDate string
	if nextDay {
		logDate = time.Now().Add(time.Second * 60).Format("2006-01-02")
	} else {
		logDate = time.Now().Format("2006-01-02")
	}
	fileName := "./logs/gs-" + logDate + ".log"
	return os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}

func checkAndRotateLogFile() {
	for {
		ticker := time.NewTicker(1 * time.Minute)
		select {
		case <-ticker.C:
			log.Println("ticker")
			mu.Lock()
			defer mu.Unlock()
			if logFile != nil {
				logFile.Close()
			}
			var err error
			logFile, err = createDailyLogFile(true)
			log.SetOutput(io.MultiWriter(logFile, os.Stdout))
			gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
			if err != nil {
				log.Fatalf("Error creating daily log file: %v", err)
			}
			mu.Unlock()
		}
	}
}
