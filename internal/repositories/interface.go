package repositories
import (
	"database/sql"
	"go-backend/internal/model"
)


type Repository interface {
	GetAll() ([]model.Task, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}