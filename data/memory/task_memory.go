package data

import (
	"errors"
	"tsn/todo/src/entities"
)

func (s *MemoryStorage) Create(task entities.Task) (*entities.Task, error) {

	task.ID = len(s.db)
	s.db = append(s.db, task)

	return &task, nil
}
func (s *MemoryStorage) Update(task entities.Task) (*entities.Task, error) {
	if len(s.db) < task.ID {
		return nil, errors.New("task not fount")
	}

	s.db[task.ID] = task

	return &task, nil
}
func (s *MemoryStorage) Delete(id int) error {
	if len(s.db)-1 < id {
		return errors.New("task not fount")
	}

	s.db = append(s.db[:id], s.db[id+1:]...)

	return nil
}
func (s *MemoryStorage) Get(id int) (entities.Task, error) {
	if len(s.db) < id {
		return entities.Task{}, errors.New("task not fount")
	}

	task := s.db[id]

	return task, nil
}

func (s *MemoryStorage) GetAll() ([]entities.Task, error) {
	return s.db, nil
}
