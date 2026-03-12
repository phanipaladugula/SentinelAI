package router

import (
	"sentinel-ai/internal/camera"
	"sentinel-ai/internal/organization"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *camera.Handler, orgHandler *organization.Handler)*gin.Engine{
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

		orgGroup := r.Group("/organizations")
	{
		orgGroup.POST("", orgHandler.Create)
		orgGroup.GET("", orgHandler.GetAll)
		orgGroup.GET("/:id", orgHandler.GetByID)
		orgGroup.DELETE("/:id", orgHandler.Delete)
	}


	return r
}