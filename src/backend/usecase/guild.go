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
	if err := name.Validate(); err != nil {

	}

	// ギルドの仮登録
	guildAggregate, err := g.guildRepository.CreateWithRegistering(ctx, name)
	if err != nil {
		return "", app.WrapError(err, app.UnexpectedFailure)
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
