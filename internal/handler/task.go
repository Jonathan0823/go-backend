package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *handler) GetAll(c *gin.Context) {
	task, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, task)
}