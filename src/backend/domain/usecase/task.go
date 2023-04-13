package usecase

import (
	"github.com/sky0621/kaubandus/domain/aggregate"
	"github.com/sky0621/kaubandus/domain/entity"
	"github.com/sky0621/kaubandus/domain/repository"
	"golang.org/x/xerrors"
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
		return xerrors.Errorf("failed to taskRepository.Save: %w", err)
	}
	return nil
}
