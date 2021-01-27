package repository

import "github.com/sky0621/kaubandus/application/aggregate"

type TaskRepository interface {
	Save(a aggregate.TaskAggregate) error
}
