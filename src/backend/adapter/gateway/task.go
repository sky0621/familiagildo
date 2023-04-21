package gateway

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func NewTaskRepository(cli *db.Client) repository.TaskRepository {
	return &taskRepository{db: cli.Q}
}

type taskRepository struct {
	db *generated.Queries
}

func (r *taskRepository) Save(a aggregate.TaskAggregate) error {

	return nil
}
