package usecase

import (
	"errors"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/repository"
)

type TaskUsecase interface {
	Add(e entity.Task) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func (u *taskUsecase) Add(e entity.Task) error {
	a := aggregate.TaskAggregate{}
	if err := u.taskRepository.Save(a); err != nil {
		return errors.Join(err)
	}
	return nil
}
