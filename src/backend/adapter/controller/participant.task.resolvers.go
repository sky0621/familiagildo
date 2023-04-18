package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
)

// CreateTaskByParticipant is the resolver for the createTaskByParticipant field.
func (r *mutationResolver) CreateTaskByParticipant(ctx context.Context, input ParticipantTaskInput) (*ParticipantTask, error) {
	panic(fmt.Errorf("not implemented: CreateTaskByParticipant - createTaskByParticipant"))
}

// UpdateTaskByParticipant is the resolver for the updateTaskByParticipant field.
func (r *mutationResolver) UpdateTaskByParticipant(ctx context.Context, input ParticipantTaskInput) (*ParticipantTask, error) {
	panic(fmt.Errorf("not implemented: UpdateTaskByParticipant - updateTaskByParticipant"))
}

// AcceptTaskByParticipant is the resolver for the acceptTaskByParticipant field.
func (r *mutationResolver) AcceptTaskByParticipant(ctx context.Context, id *string) (*bool, error) {
	panic(fmt.Errorf("not implemented: AcceptTaskByParticipant - acceptTaskByParticipant"))
}

// ListTaskByParticipant is the resolver for the listTaskByParticipant field.
func (r *queryResolver) ListTaskByParticipant(ctx context.Context) ([]*ParticipantTask, error) {
	panic(fmt.Errorf("not implemented: ListTaskByParticipant - listTaskByParticipant"))
}

// FindTaskByParticipant is the resolver for the findTaskByParticipant field.
func (r *queryResolver) FindTaskByParticipant(ctx context.Context, filter *ParticipantTaskFilter) ([]*ParticipantTask, error) {
	panic(fmt.Errorf("not implemented: FindTaskByParticipant - findTaskByParticipant"))
}

// GetTaskByParticipant is the resolver for the getTaskByParticipant field.
func (r *queryResolver) GetTaskByParticipant(ctx context.Context, id string) (*ParticipantTask, error) {
	panic(fmt.Errorf("not implemented: GetTaskByParticipant - getTaskByParticipant"))
}