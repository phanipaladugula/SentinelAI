package main

import (
	"net/http"
	"sentinel-ai/internal/config"
	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	cfg:=config.LoadConfig()

	r.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
		})
	})

	r.Run(":"+cfg.Port)
}