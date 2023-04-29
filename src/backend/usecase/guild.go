package usecase

import (
	"context"
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

func NewGuild(r repository.GuildRepository) GuildInputPort {
	return &guildInteractor{guildRepository: r}
}

type guildInteractor struct {
	guildRepository repository.GuildRepository
}

// RequestCreateGuildByGuest is ギルド登録を依頼して受付番号を返す
func (g *guildInteractor) RequestCreateGuildByGuest(ctx context.Context, name vo.GuildName, mail vo.OwnerMail) (string, error) {
	{
		var customErrors app.CustomErrors
		for _, v := range []vo.ValueObject[string]{name, mail} {
			if err := v.Validate(); err != nil {
				customErrors = append(customErrors, app.NewCustomError(
					err, app.ValidationFailure, app.NewCustomErrorDetail(v.FieldName(), v.ToVal())))
			}
		}
		if len(customErrors) > 0 {
			return "", customErrors
		}
	}

	// ギルドの仮登録
	guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, name)
	if err != nil {
		return "", app.NewCustomError(err, app.UnexpectedFailure, nil)
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
