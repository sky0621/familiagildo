package usecase

import (
	"context"

	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/event"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildInputPort interface {
	RequestCreateGuildByGuest(ctx context.Context, input RequestCreateGuildInput) (vo.AcceptedNumber, error)
	CreateGuildByGuest(ctx context.Context, input CreateGuildByGuestInput) error
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
	customErrors := service.Validate(ctx, []vo.ValueObject[string]{input.GuildName, input.OwnerMail})
	if len(customErrors) > 0 {
		return "", customErrors
	}

	{
		validToken, err := g.guestTokenRepository.GetByOwnerMailWithinValidPeriod(ctx, input.OwnerMail)
		if err != nil {
			return "", app.NewUnexpectedError(ctx, err)
		}

		if validToken != nil {
			r := validToken.Root
			if r == nil {
				return "", app.NewUnexpectedError(ctx, app.NewError("validToken.Root is nil"))
			}
			return "", app.NewAlreadyExistsError(ctx, r.Token.FieldName(), r.Token.ToVal())
		}
	}

	acceptedNumber := service.CreateAcceptedNumber()

	if err := g.transactionRepository.ExecInTransaction(ctx, func(ctx context.Context) error {
		guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, input.GuildName)
		if err != nil {
			return app.WithStackError(err)
		}

		token := service.CreateToken()

		expirationDate := service.CreateGuestTokenExpirationDate()

		_, err = g.guestTokenRepository.Create(ctx, guildAggregate.Root.ID, input.OwnerMail,
			&entity.GuestToken{Token: token, ExpirationDate: expirationDate},
			acceptedNumber)
		if err != nil {
			return app.WithStackError(err)
		}

		if err := g.guildEvent.CreateRequested(ctx, event.CreateRequestedInput{
			GuildName:      input.GuildName,
			Token:          token,
			ExpirationDate: expirationDate,
			OwnerMail:      input.OwnerMail,
			AcceptedNumber: acceptedNumber,
		}); err != nil {
			return app.WithStackError(err)
		}

		return nil
	}); err != nil {
		return "", app.NewUnexpectedError(ctx, err)
	}

	return acceptedNumber, nil
}

// GetGuildByToken is トークンに紐づくギルド情報を返す
func (g *guildInteractor) GetGuildByToken(ctx context.Context, token vo.Token) (*aggregate.Guild, error) {
	customErrors := service.Validate(ctx, []vo.ValueObject[string]{token})
	if len(customErrors) > 0 {
		return nil, customErrors
	}

	validToken, err := g.guestTokenRepository.GetByTokenWithinValidPeriod(ctx, token)
	if err != nil {
		return nil, app.NewUnexpectedError(ctx, err)
	}

	if validToken == nil || validToken.Root == nil || validToken.Guild == nil {
		return nil, app.NewAuthorizationError(ctx, token.FieldName(), token.ToVal())
	}

	guild, err := g.guildRepository.GetByID(ctx, validToken.Guild.ID)
	if err != nil {
		return nil, app.NewUnexpectedError(ctx, err)
	}

	if guild == nil {
		return nil, app.NewUnexpectedError(ctx, app.NewError("guild is nil"))
	}

	guild.Owner = &entity.Owner{Mail: validToken.Root.Mail}

	return guild, nil
}

func (g *guildInteractor) CreateGuildByGuest(ctx context.Context, input CreateGuildByGuestInput) error {
	customErrors := service.Validate(ctx, []vo.ValueObject[string]{input.Token, input.OwnerName, input.LoginID, input.Password})
	if len(customErrors) > 0 {
		return customErrors
	}

	// FIXME:
	return nil
}
