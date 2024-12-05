package handler

import (
	"go-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}