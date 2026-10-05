package main

import (
	"log"
	"quick/internal/config"
	"quick/internal/health"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}

	router := gin.Default()

	api := router.Group("/socket/v1")
	health.RegisterRoutes(api)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
