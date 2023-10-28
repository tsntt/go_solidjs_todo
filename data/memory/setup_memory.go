package data

import (
	"tsn/todo/src/entities"
)

type MemoryStorage struct {
	db []entities.Task
}

func NewMemoryStorage(db []entities.Task) *MemoryStorage {
	return &MemoryStorage{
		db: db,
	}
}
