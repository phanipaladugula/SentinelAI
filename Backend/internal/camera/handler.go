package camera

import "github.com/gin-gonic/gin"

type Handler struct{
	service *Service
}

func NewHandler(s *Service)*Handler{
	return &Handler{service:s}
}

func(h *Handler)Create(c *gin.Context){
	var req CreateCameraRequest

	if err:=c.ShouldBindBodyWithJSON(&req);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
	}

	res,err:=h.service.Create(req)
	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(201,res)
}

func(h *Handler)GetAll(c *gin.Context){
	res,err:=h.service.GetAll()

	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,res)
}

func(h *Handler)GetByID(c *gin.Context){
	id:=c.Param("id")
	res,err:=h.service.GetByID(id)
	if err!=nil{
		c.JSON(404,gin.H{"error":"not found"})
		return
	}
	c.JSON(200,res)
}

func(h *Handler)Delete(c *gin.Context){
	id:=c.Param("id")
	err:=h.service.Delete(id)
	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{"message":"deleted"})
}