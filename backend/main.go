package main

import (
	"log"
	"time"

	"backend/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	repo, err := repository.InitDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 日志查询API
	r.GET("/api/logs", repo.GetLogs)

	log.Println("服务器启动在 :8080 端口")
	r.Run(":8080")
}



