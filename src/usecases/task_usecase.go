package usecases

import (
	"time"
	"tsn/todo/src/entities"
	"tsn/todo/src/util"
)

type TaskInteractor struct {
	TaskInteractor entities.TaskRepository
}

func NewTaskInteractor(store entities.TaskRepository) *TaskInteractor {
	return &TaskInteractor{
		TaskInteractor: store,
	}
}

func (interactor *TaskInteractor) Create(content, description, due string) (*entities.Task, error) {
	timeDue := util.StringToTimeUnix(due)

	newTask := entities.Task{
		Content:     content,
		Description: description,
		Status:      false,
		Due:         timeDue,
		CreatedAt:   time.Now().UTC(),
	}

	createdTask, err := interactor.TaskInteractor.Create(newTask)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (interactor *TaskInteractor) Update(id int, content, description, due string) (*entities.Task, error) {
	timeDue := util.StringToTimeUnix(due)

	task, err := interactor.TaskInteractor.Get(id)
	if err != nil {
		return nil, err
	}

	task.Content = content
	task.Description = description
	task.Due = timeDue

	updatedTask, err := interactor.TaskInteractor.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (interactor *TaskInteractor) ChangeStatus(id int) (*entities.Task, error) {
	task, err := interactor.TaskInteractor.Get(id)
	if err != nil {
		return nil, err
	}

	task.ChangeStatus()

	updatedTask, err := interactor.TaskInteractor.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (interactor *TaskInteractor) Delete(id int) error {
	err := interactor.TaskInteractor.Delete(id)

	return err
}

func (interactor *TaskInteractor) Get(id int) (entities.Task, error) {
	task, err := interactor.TaskInteractor.Get(id)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (interactor *TaskInteractor) GetAll() ([]entities.Task, error) {
	tasks, err := interactor.TaskInteractor.GetAll()
	if err != nil {
		return []entities.Task{}, err
	}

	return tasks, nil
}
