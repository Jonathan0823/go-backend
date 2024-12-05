package handler

import (
	"go-backend/internal/model"
	"go-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *handler {
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

func (h *handler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *handler) DeleteTask(c *gin.Context){
	id := c.Query("id")

	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}