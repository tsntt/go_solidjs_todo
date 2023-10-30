package data

import (
	"fmt"
	"tsn/todo/src/entities"
)

func (s *PostgresStorage) Create(t entities.Task) (*entities.Task, error) {
	q := "INSERT INTO tasks( content, description, due, created_at ) VALUES ($1, $2, $3, $4) RETURNING *"

	task := entities.Task{}

	err := s.db.QueryRow(q, t.Content, t.Description, t.Due, t.CreatedAt).Scan(
		&task.ID,
		&task.Content,
		&task.Description,
		&task.Status,
		&task.Due,
		&task.CreatedAt,
	)

	fmt.Print(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *PostgresStorage) Update(task entities.Task) (*entities.Task, error) {
	q := "UPDATE tasks SET (content, description, status, due) = ($2, $3, $4, $5) WHERE id=$1 RETURNING *"

	row := s.db.QueryRow(q, task.ID, task.Content, task.Description, task.Status, task.Due)

	t := entities.Task{}

	if err := row.Scan(
		&t.ID,
		&t.Content,
		&t.Description,
		&t.Status,
		&t.Due,
		&t.CreatedAt); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return &t, nil
}
func (s *PostgresStorage) Delete(id int) error {
	q := "DELETE FROM tasks WHERE id=$1"

	_, err := s.db.Exec(q, id)

	return err
}

func (s *PostgresStorage) Get(id int) (entities.Task, error) {
	q := "SELECT * FROM tasks WHERE id=$1"

	task := entities.Task{}
	res := s.db.QueryRow(q, id)
	if err := res.Scan(
		&task.ID,
		&task.Content,
		&task.Description,
		&task.Status,
		&task.Due,
		&task.CreatedAt); err != nil {
		return task, err
	}

	return task, nil
}

func (s *PostgresStorage) GetAll() ([]entities.Task, error) {
	q := "SELECT * FROM tasks ORDER BY id"

	tasks := []entities.Task{}

	rows, err := s.db.Query(q)
	if err != nil {
		return []entities.Task{}, err
	}

	for rows.Next() {
		task := entities.Task{}

		if err := rows.Scan(&task.ID,
			&task.Content,
			&task.Description,
			&task.Status,
			&task.Due,
			&task.CreatedAt); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}
