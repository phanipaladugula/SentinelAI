package main

import (
	"net/http"

	"go.uber.org/zap"

	"sentinel-ai/internal/app"
	"sentinel-ai/internal/camera"
	"sentinel-ai/internal/config"
	"sentinel-ai/internal/logger"
	"sentinel-ai/internal/router"
)

func main(){

	cfg:=config.LoadConfig()

	logger.Init()

	a,err:=app.NewDB(cfg)
	if err!=nil{
		logger.Log.Fatal("app intialization failed",zap.Error(err))
	}

	r:=a.Router

	r.Run(":"+cfg.Port)
}