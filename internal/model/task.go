package model

type Task struct {
	ID          int    `json:"id"`
	Title 	 	string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate	 	string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}