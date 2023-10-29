package entities

import "time"

type Task struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Due         time.Time `json:"due"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository interface {
	Create(task Task) (*Task, error)
	Update(task Task) (*Task, error)
	Delete(id int) error
	Get(id int) (Task, error)
	GetAll() ([]Task, error)
}

func (t *Task) ChangeStatus() {
	t.Status = !t.Status
}
