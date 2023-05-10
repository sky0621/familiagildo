package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"errors"
	"fmt"

	"github.com/sky0621/familiagildo/adapter/controller/custommodel"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/app/log"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/usecase"
)

// RequestCreateGuildByGuest is the resolver for the requestCreateGuildByGuest field.
func (r *mutationResolver) RequestCreateGuildByGuest(ctx context.Context, input RequestCreateGuildInput) (*GuestToken, error) {
	usecaseInput := usecase.RequestCreateGuildInput{
		GuildName: vo.ToGuildName(input.GuildName),
		OwnerMail: vo.ToOwnerMail(input.OwnerMail),
	}

	acceptedNumber, err := r.GuildUsecase.RequestCreateGuildByGuest(ctx, usecaseInput)
	if err != nil {
		var cErr *app.CustomError
		if errors.As(err, &cErr) && cErr.GetErrorCode() == app.AlreadyExistsError {
			log.Warn(cErr.Error())
		} else {
			log.ErrorSend(err)
		}
		return nil, CreateGQLError(ctx, err)
	}

	return &GuestToken{
		AcceptedNumber: acceptedNumber.ToVal(),
	}, err
}

// GetGuildByToken is the resolver for the getGuildByToken field.
func (r *queryResolver) GetGuildByToken(ctx context.Context, token string) (*Guild, error) {
	usecaseToken := vo.ToToken(token)

	guild, err := r.GuildUsecase.GetGuildByToken(ctx, usecaseToken)
	if err != nil {
		log.ErrorSend(err)
		return nil, CreateGQLError(ctx, err)
	}
	if guild.Root == nil {
		// FIXME:
		return nil, nil
	}
	if guild.Owner == nil {
		// FIXME:
		return nil, nil
	}

	result := &Guild{
		Name: guild.Root.Name.ToVal(),
		// FIXME:
		Owner: &Owner{
			Mail: guild.Owner.Mail.ToVal(),
		},
	}
	return result, nil
}

// CreateGuildByGuest is the resolver for the createGuildByGuest field.
func (r *mutationResolver) CreateGuildByGuest(ctx context.Context, input CreateGuildByGuestInput) (*custommodel.Void, error) {
	usecaseInput := usecase.CreateGuildByGuestInput{
		Token:     vo.ToToken(input.Token),
		OwnerName: vo.ToOwnerName(input.OwnerName),
		LoginID:   vo.ToLoginID(input.LoginID),
		Password:  vo.ToPassword(input.Password),
	}

	if err := r.GuildUsecase.CreateGuildByGuest(ctx, usecaseInput); err != nil {
		log.ErrorSend(err)
		return nil, CreateGQLError(ctx, err)
	}

	return &custommodel.Void{}, nil
}

// CreateParticipantByGuest is the resolver for the createParticipantByGuest field.
func (r *mutationResolver) CreateParticipantByGuest(ctx context.Context, input CreateParticipantByGuestInput) (*custommodel.Void, error) {
	panic(fmt.Errorf("not implemented: CreateParticipantByGuest - createParticipantByGuest"))
}
