package main

import (
	"go.uber.org/zap"
	"sentinel-ai/internal/app"
	"sentinel-ai/internal/config"
	"sentinel-ai/internal/logger"
)

func main() {

	cfg := config.LoadConfig()

	logger.Init()

	a, err := app.New(cfg) 
	if err != nil {
		logger.Log.Fatal("app initialization failed", zap.Error(err))
	}

	a.Router.Run(":" + cfg.Port)
}