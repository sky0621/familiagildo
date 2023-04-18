package domain

import "github.com/sky0621/familiagildo/domain/entity"

type Task interface {
	Create(e entity.Task) (entity.Task, error)
	Tasks() ([]entity.Task, error)
}

func NewTask() Task {
	return &task{}
}

type task struct {
}

func (t *task) Create(e entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (t *task) Tasks() ([]entity.Task, error) {
	return []entity.Task{}, nil
}
