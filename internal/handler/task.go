package handler

import (
	"go-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
}

type handler struct {
	service service.Service
}

func TaskHandler(service service.Service) *handler {
	return &handler{service}
}

func (h *handler) GetAll(c *gin.Context) {
	task, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, task)
}