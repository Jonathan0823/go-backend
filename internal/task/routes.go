package task

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine ,db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)


	r.Group("/tasks")
	{
		r.GET("/", handler.GetAll)
	}

}