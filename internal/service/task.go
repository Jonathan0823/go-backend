package service

import (
	"go-backend/internal/model"
)

func (s *service) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}