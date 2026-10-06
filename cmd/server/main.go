package main

import (
	"log"
	"quick/internal/config"
	"quick/internal/health"

	"github.com/gin-gonic/gin"
)

// @title Quick Chat Realtime API
// @version 1.0
// @description Quick Chat Realtime Server API
// @BasePath /realtime/v1
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}

	switch cfg.Env {
	case config.EnvDev, config.EnvTest:
		gin.SetMode(gin.DebugMode)

	case config.EnvStage, config.EnvProd:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	api := router.Group("/realtime/v1")
	health.RegisterRoutes(api)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
