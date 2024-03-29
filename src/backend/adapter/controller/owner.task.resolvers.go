package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"github.com/sky0621/familiagildo/adapter/controller/model"
)

// CreateTaskByOwner is the resolver for the createTaskByOwner field.
func (r *mutationResolver) CreateTaskByOwner(ctx context.Context, input model.OwnerTaskInput) (*model.OwnerTask, error) {
	// FIXME: middleware で ctx に積んだ user_id, role 等から、認証チェック・認可チェックを行う！（このHandlerを実行してよいか否かのチェックはHandlerの責務）

	panic(fmt.Errorf("not implemented: CreateTaskByOwner - createTaskByOwner"))
}

// UpdateTaskByOwner is the resolver for the updateTaskByOwner field.
func (r *mutationResolver) UpdateTaskByOwner(ctx context.Context, input model.OwnerTaskInput) (*model.OwnerTask, error) {
	panic(fmt.Errorf("not implemented: UpdateTaskByOwner - updateTaskByOwner"))
}

// DeleteTaskByOwner is the resolver for the deleteTaskByOwner field.
func (r *mutationResolver) DeleteTaskByOwner(ctx context.Context, input model.OwnerTaskInput) (*string, error) {
	panic(fmt.Errorf("not implemented: DeleteTaskByOwner - deleteTaskByOwner"))
}

// AcceptTaskByOwner is the resolver for the acceptTaskByOwner field.
func (r *mutationResolver) AcceptTaskByOwner(ctx context.Context, id *string) (*bool, error) {
	panic(fmt.Errorf("not implemented: AcceptTaskByOwner - acceptTaskByOwner"))
}

// ListTaskByOwner is the resolver for the listTaskByOwner field.
func (r *queryResolver) ListTaskByOwner(ctx context.Context) ([]*model.OwnerTask, error) {
	panic(fmt.Errorf("not implemented: ListTaskByOwner - listTaskByOwner"))
}

// FindTaskByOwner is the resolver for the findTaskByOwner field.
func (r *queryResolver) FindTaskByOwner(ctx context.Context, filter *model.OwnerTaskFilter) ([]*model.OwnerTask, error) {
	panic(fmt.Errorf("not implemented: FindTaskByOwner - findTaskByOwner"))
}

// GetTaskByOwner is the resolver for the getTaskByOwner field.
func (r *queryResolver) GetTaskByOwner(ctx context.Context, id string) (*model.OwnerTask, error) {
	panic(fmt.Errorf("not implemented: GetTaskByOwner - getTaskByOwner"))
}
