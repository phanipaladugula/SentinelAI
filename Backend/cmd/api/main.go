package main

import (
	"fmt"
	"go.uber.org/zap"
	"sentinel-ai/internal/app"
	"sentinel-ai/internal/camera/config"
	"sentinel-ai/internal/logger"
)

func main() {

	cfg := config.LoadConfig()
	fmt.Printf("DEBUG: Host='%s' Port='%s' User='%s' Pass='%s'\n", 
    cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword)

	logger.Init()

	a, err := app.New(cfg) 
	if err != nil {
		logger.Log.Fatal("app initialization failed", zap.Error(err))
	}

	a.Router.Run(":" + cfg.Port)
}