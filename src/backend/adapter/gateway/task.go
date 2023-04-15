package gateway

import (
	"github.com/sky0621/kaubandus/domain/aggregate"
	"github.com/sky0621/kaubandus/domain/repository"
	"github.com/sky0621/kaubandus/external/db"
)

func NewTaskRepository(db *db.Queries) repository.TaskRepository {
	return &taskRepository{db: db}
}

type taskRepository struct {
	db *db.Queries
}

func (r *taskRepository) Save(a aggregate.TaskAggregate) error {

	return nil
}
