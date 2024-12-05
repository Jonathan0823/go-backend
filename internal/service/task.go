package service

import (
	"go-backend/internal/model"
	"go-backend/internal/repositories"
)

type TaskService interface {
	GetAll() ([]model.Task, error)
	CreateTask(task model.Task) error
	DeleteTask(id string) error
}

type service struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *service) CreateTask(task model.Task) error {
	return s.repo.CreateTask(task)
}

func (s *service) DeleteTask(id string) error{
	return s.repo.DeleteTask(id)
}