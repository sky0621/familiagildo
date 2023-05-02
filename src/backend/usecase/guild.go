package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildInputPort interface {
	// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
	RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error)
}

func NewGuild(tr repository.TransactionRepository, gtr repository.GuestTokenRepository, gr repository.GuildRepository) GuildInputPort {
	return &guildInteractor{transactionRepository: tr, guestTokenRepository: gtr, guildRepository: gr}
}

type guildInteractor struct {
	transactionRepository repository.TransactionRepository
	guestTokenRepository  repository.GuestTokenRepository
	guildRepository       repository.GuildRepository
}

// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
func (g *guildInteractor) RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error) {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{name, mail} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewValidationError(err, v.FieldName(), v.ToVal()))
			}
		}
		if len(customErrors) > 0 {
			return "", customErrors
		}
	}

	{
		validToken, err := g.guestTokenRepository.GetByOwnerMailWithinValidPeriod(ctx, mail)
		if err != nil {
			return "", app.NewUnexpectedError(err)
		}

		if validToken != nil {
			r := validToken.Root
			if r == nil {
				return "", app.NewUnexpectedError(errors.New("validToken.Root is nil"))
			}
			return "", app.NewAlreadyExistsError(r.Token.FieldName(), r.Token.ToVal())
		}
	}

	// トークンの生成
	token := service.CreateToken()

	// 有効期限も生成
	expirationDate := service.CreateGuestTokenExpirationDate()

	// 受付番号の生成
	acceptedNumber := service.CreateAcceptedNumber()

	if err := g.transactionRepository.ExecInTransaction(ctx, func(ctx context.Context) error {
		guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, name)
		if err != nil {
			return errors.WithStack(err)
		}

		_, err = g.guestTokenRepository.Create(ctx, guildAggregate.Root.ID, mail,
			&entity.GuestToken{Token: vo.ParseToken(token), ExpirationDate: vo.ParseExpirationDate(expirationDate)},
			vo.ParseAcceptedNumber(acceptedNumber))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return "", app.NewUnexpectedError(err)
	}

	// FIXME: メール送信

	return acceptedNumber, nil
}
