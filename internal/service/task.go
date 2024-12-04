package service

import (
	"go-backend/internal/model"
	"go-backend/internal/repositories"
)

type Service interface {
	GetAll() ([]model.Task, error)
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}