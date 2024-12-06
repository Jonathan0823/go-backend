package repositories

import (
	"database/sql"
	"go-backend/internal/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	CreateUser(user model.User) error
	DeleteUser(id string) error
	EditUser(user model.User) error
	GetUserByID(id string) (model.User, error)
}

type userrepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userrepository {
	return &userrepository{db}
}

func (r *userrepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userrepository) CreateUser(user model.User) error {
	_, err := r.db.Exec(`INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *userrepository) DeleteUser(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userrepository) EditUser(user model.User) error {
	_, err := r.db.Exec(`UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4`, user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userrepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}