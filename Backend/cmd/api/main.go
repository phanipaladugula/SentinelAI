package main

import (
	"net/http"
	"sentinel-ai/internal/config"
	"github.com/gin-gonic/gin"
	"sentinel-ai/internal/logger"
)

func main(){
	r:=gin.Default()
	cfg:=config.LoadConfig()
	logger.Init()


	r.GET("/health",func(c *gin.Context){

		logger.Log.Info("health check called")
		
		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
		})
	})

	r.Run(":"+cfg.Port)
}