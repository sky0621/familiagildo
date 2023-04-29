package repository

import (
	"context"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuestTokenRepository interface {
	// GetByOwnerMailWithinValidPeriod is 指定のメールアドレスで登録済みの有効期限内のトークンを返す
	GetByOwnerMailWithinValidPeriod(ctx context.Context, mail vo.OwnerMail) (*aggregate.GuestToken, error)
}
