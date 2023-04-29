package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildInputPort interface {
	// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
	RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error)
}

func NewGuild(tr repository.GuestTokenRepository, gr repository.GuildRepository) GuildInputPort {
	return &guildInteractor{tokenRepository: tr, guildRepository: gr}
}

type guildInteractor struct {
	tokenRepository repository.GuestTokenRepository
	guildRepository repository.GuildRepository
}

// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
func (g *guildInteractor) RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error) {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{name, mail} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewCustomError(
					err, app.ValidationError, app.NewCustomErrorDetail(v.FieldName(), v.ToVal())))
			}
		}
		if len(customErrors) > 0 {
			return "", customErrors
		}
	}

	validToken, err := g.tokenRepository.GetByOwnerMailWithinValidPeriod(ctx, mail)
	if err != nil {
		return "", app.NewCustomError(err, app.UnexpectedError, nil)
	}

	if validToken != nil {
		r := validToken.Root
		if r == nil {
			return "", app.NewCustomError(errors.New("validToken.Root is nil"), app.UnexpectedError, nil)
		}
		return "", app.NewCustomError(nil, app.AlreadyExistsError, app.NewCustomErrorDetail(r.Token.FieldName(), r.Token.ToVal()))
	}

	guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, name)
	if err != nil {
		return "", app.NewCustomError(err, app.UnexpectedError, nil)
	}
	// FIXME:
	fmt.Println(guildAggregate)

	// トークンの生成
	// FIXME: 有効期限も生成
	token := service.CreateToken()
	// FIXME:
	fmt.Println(token)

	// トークンの保存

	// メール送信

	// 受付番号の生成
	acceptedNumber := service.CreateAcceptNumber()

	return acceptedNumber, nil
}
