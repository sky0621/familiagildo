package application

import "github.com/sky0621/kaubandus/application/entity"

type Task interface {
	Create() (entity.Task, error)
	Tasks() ([]entity.Task, error)
}

func NewTask() Task {
	return &task{}
}

type task struct {
}

func (t *task) Create() (entity.Task, error) {
	return entity.Task{}, nil
}

func (t *task) Tasks() ([]entity.Task, error) {
	return []entity.Task{}, nil
}
