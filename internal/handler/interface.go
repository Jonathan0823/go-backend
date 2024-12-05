package handler

import (
	"go-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}