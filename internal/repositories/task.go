package repositories

import (
	"database/sql"
	"go-backend/internal/model"
)

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	CreateTask(task model.Task) error
	DeleteTask(id string) error
	EditStatus(id string, status string) error
	EditTask(task model.Task) error
}

type taskrepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskrepository {
	return &taskrepository{db}
}

func (r *taskrepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskrepository) CreateTask(task model.Task) error {
	_, err := r.db.Exec(`INSERT INTO tasks (title, description, status, due_date) VALUES ($1, $2, $3, $4)`, task.Title, task.Description, task.Status, task.DueDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskrepository) DeleteTask(id string) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskrepository) EditStatus(id string, status string) error {
	_, err := r.db.Exec(`UPDATE tasks SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskrepository) EditTask(task model.Task) error {
	_, err := r.db.Exec(`UPDATE tasks SET 
								title = $1,
								description = $2,
								status = $3,
								due_date = $4
							WHERE id = $5`,
		task.Title, task.Description, task.Status, task.DueDate, task.ID)
	if err != nil {
		return err
	}
	return nil

}
