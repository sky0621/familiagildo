package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/event"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildInputPort interface {
	// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
	RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error)
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
func (g *guildInteractor) RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error) {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{name, mail} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewValidationError(ctx, err, v.FieldName(), v.ToVal()))
			}
		}
		if len(customErrors) > 0 {
			return "", customErrors
		}
	}

	{
		validToken, err := g.guestTokenRepository.GetByOwnerMailWithinValidPeriod(ctx, mail)
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
		guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, name)
		if err != nil {
			return errors.WithStack(err)
		}

		token := service.CreateToken()

		expirationDate := service.CreateGuestTokenExpirationDate()

		_, err = g.guestTokenRepository.Create(ctx, guildAggregate.Root.ID, mail,
			&entity.GuestToken{Token: token, ExpirationDate: expirationDate},
			acceptedNumber)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := g.guildEvent.CreateRequested(ctx, event.CreateRequestedInput{
			GuildName:      name,
			Token:          token,
			ExpirationDate: expirationDate,
			OwnerMail:      mail,
			AcceptedNumber: acceptedNumber,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return "", app.NewUnexpectedError(ctx, err)
	}

	return acceptedNumber.ToVal(), nil
}
