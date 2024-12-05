package service

import (
	"go-backend/internal/model"
	"go-backend/internal/repositories"
)

type Service interface {
	GetAll() ([]model.Task, error)
	CreateTask(task model.Task) error
	DeleteTask(id string) error
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) *service {
	return &service{repo}
}