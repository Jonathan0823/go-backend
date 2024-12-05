package repositories
import (
	"database/sql"
	"go-backend/internal/model"
)


type Repository interface {
	GetAll() ([]model.Task, error)
	CreateTask(task model.Task) error
	DeleteTask(id string) error
	EditStatus(id string, status string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}