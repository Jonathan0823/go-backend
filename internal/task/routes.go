package task

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	r := gin.Default()

	r.Group("/tasks")
	{
		r.GET("/", handler.GetAll)
	}

}