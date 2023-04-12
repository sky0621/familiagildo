package web

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/kaubandus/adapter/web/gqlmodel"
)

func (r *mutationResolver) CreateTask(ctx context.Context, task gqlmodel.TaskInput) (*gqlmodel.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindTask(ctx context.Context) ([]*gqlmodel.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
