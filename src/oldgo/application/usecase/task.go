package usecase

import (
	"github.com/sky0621/kaubandus/application/entity"
	"github.com/sky0621/kaubandus/application/repository"
	"golang.org/x/xerrors"
)

type TaskUsecase interface {
	Add(e entity.Task) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func (u *taskUsecase) Add(e entity.Task) error {
	if err := u.taskRepository.Save(); err != nil {
		return xerrors.Errorf("failed to taskRepository.Save: %w", err)
	}
	return nil
}
