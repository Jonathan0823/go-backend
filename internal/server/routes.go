package server

import (
	"go-backend/internal/repositories"
	"go-backend/internal/service"
	"go-backend/internal/handler"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()


	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	repo := repositories.NewRepository(s.db.GetDB())
	svc := service.NewService(repo)
	mainHandler := handler.NewHandler(svc)

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/", mainHandler.GetAll)
		taskGroup.POST("/", mainHandler.CreateTask)
		taskGroup.DELETE("/", mainHandler.DeleteTask)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
