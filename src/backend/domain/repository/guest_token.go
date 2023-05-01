package repository

import (
	"context"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuestTokenRepository interface {
	// GetByOwnerMailWithinValidPeriod is 指定のメールアドレスで登録済みの有効期限内のトークンを返す
	GetByOwnerMailWithinValidPeriod(ctx context.Context, mail vo.OwnerMail) (*aggregate.GuestToken, error)
	// Create is トークンの登録
	Create(ctx context.Context, guildID vo.ID, mail vo.OwnerMail, guestToken *entity.GuestToken, acceptedNumber vo.AcceptedNumber) (*aggregate.GuestToken, error)
}
