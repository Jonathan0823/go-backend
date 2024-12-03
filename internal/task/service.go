package task

type Service interface {
	GetAll() ([]Task, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]Task, error) {
	return s.repo.getAll()
}