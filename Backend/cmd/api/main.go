package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"sentinel-ai/internal/config"
	"sentinel-ai/internal/database"
	"sentinel-ai/internal/logger"
)

func main(){
	r:=gin.Default()
	cfg:=config.LoadConfig()
	logger.Init()
	db,err := database.NewDB(cfg)


	if err != nil{
		logger.Log.Fatal("database connection failed",zap.Error(err))
	}
	
	defer db.Close()
	r.GET("/health",func(c *gin.Context){

		logger.Log.Info("health check called")

		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
		})
	})

	r.Run(":"+cfg.Port)
}