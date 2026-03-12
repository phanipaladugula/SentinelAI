package router

import (
	"sentinel-ai/internal/camera"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *camera.Handler)*gin.Engine{
	r:=gin.Default()

	r.GET("/health",func(c *gin.Context){
		c.JSON(200,gin.H{"status":"ok"})
	})

	cameraGroup:=r.Group("/cameras")
	{
		cameraGroup.POST("",h.Create)
		cameraGroup.GET("",h.GetAll)
		cameraGroup.GET("/:id",h.GetByID)
		cameraGroup.DELETE("/:id",h.Delete)
	}
	return r
}