package camera

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Handler struct {
    service *Service
}

func NewHandler(s *Service) *Handler {
    return &Handler{service: s}
}

func (h *Handler) Create(c *gin.Context) {
    var req CreateCameraRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Passing Request Context
    res, err := h.service.Create(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetAll(c *gin.Context) {
    res, err := h.service.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}

func (h *Handler) GetByID(c *gin.Context) {
    id := c.Param("id")
    res, err := h.service.GetByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "camera not found"})
        return
    }
    c.JSON(http.StatusOK, res)
}

func (h *Handler) Delete(c *gin.Context) {
    id := c.Param("id")
    err := h.service.Delete(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}