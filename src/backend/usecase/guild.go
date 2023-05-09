package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/event"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildInputPort interface {
	// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
	RequestCreateGuildByGuest(ctx context.Context, input RequestCreateGuildInput) (vo.AcceptedNumber, error)
	// CreateGuildByGuest is ギルド及びオーナー情報を登録する
	CreateGuildByGuest(ctx context.Context, input CreateGuildByGuestInput) error
	// GetGuildByToken is
	GetGuildByToken(ctx context.Context, token vo.Token) (*aggregate.Guild, error)
}

type RequestCreateGuildInput struct {
	GuildName vo.GuildName
	OwnerMail vo.OwnerMail
}

type CreateGuildByGuestInput struct {
	Token     vo.Token
	OwnerName vo.OwnerName
	LoginID   vo.LoginID
	Password  vo.Password
}

func NewGuild(
	tr repository.TransactionRepository,
	gtr repository.GuestTokenRepository,
	gr repository.GuildRepository,
	ge event.GuildEvent,
) GuildInputPort {
	return &guildInteractor{transactionRepository: tr, guestTokenRepository: gtr, guildRepository: gr, guildEvent: ge}
}

type guildInteractor struct {
	transactionRepository repository.TransactionRepository
	guestTokenRepository  repository.GuestTokenRepository
	guildRepository       repository.GuildRepository

	guildEvent event.GuildEvent
}

// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
func (g *guildInteractor) RequestCreateGuildByGuest(ctx context.Context, input RequestCreateGuildInput) (vo.AcceptedNumber, error) {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{input.GuildName, input.OwnerMail} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewValidationError(ctx, err, v.FieldName(), v.ToVal()))
			}
		}
		if len(customErrors) > 0 {
			return "", customErrors
		}
	}

	{
		validToken, err := g.guestTokenRepository.GetByOwnerMailWithinValidPeriod(ctx, input.OwnerMail)
		if err != nil {
			return "", app.NewUnexpectedError(ctx, err)
		}

		if validToken != nil {
			r := validToken.Root
			if r == nil {
				return "", app.NewUnexpectedError(ctx, errors.New("validToken.Root is nil"))
			}
			return "", app.NewAlreadyExistsError(ctx, r.Token.FieldName(), r.Token.ToVal())
		}
	}

	acceptedNumber := service.CreateAcceptedNumber()

	if err := g.transactionRepository.ExecInTransaction(ctx, func(ctx context.Context) error {
		guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, input.GuildName)
		if err != nil {
			return errors.WithStack(err)
		}

		token := service.CreateToken()

		expirationDate := service.CreateGuestTokenExpirationDate()

		_, err = g.guestTokenRepository.Create(ctx, guildAggregate.Root.ID, input.OwnerMail,
			&entity.GuestToken{Token: token, ExpirationDate: expirationDate},
			acceptedNumber)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := g.guildEvent.CreateRequested(ctx, event.CreateRequestedInput{
			GuildName:      input.GuildName,
			Token:          token,
			ExpirationDate: expirationDate,
			OwnerMail:      input.OwnerMail,
			AcceptedNumber: acceptedNumber,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return "", app.NewUnexpectedError(ctx, err)
	}

	return acceptedNumber, nil
}

func (g *guildInteractor) GetGuildByToken(ctx context.Context, token vo.Token) (*aggregate.Guild, error) {

	// FIXME:
	return nil, nil
}

func (g *guildInteractor) CreateGuildByGuest(ctx context.Context, input CreateGuildByGuestInput) error {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{input.Token, input.OwnerName, input.LoginID, input.Password} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewValidationError(ctx, err, v.FieldName(), v.ToVal()))
			}
		}
		if len(customErrors) > 0 {
			return customErrors
		}
	}

	// FIXME:
	return nil
}
