package repository

import (
	"context"

	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuestTokenRepository interface {
	GetByOwnerMailWithinValidPeriod(ctx context.Context, mail vo.OwnerMail) (*aggregate.GuestToken, error)
	GetByTokenWithinValidPeriod(ctx context.Context, token vo.Token) (*aggregate.GuestToken, error)
	Create(ctx context.Context, guildID vo.ID, mail vo.OwnerMail, guestToken *entity.GuestToken, acceptedNumber vo.AcceptedNumber) (*aggregate.GuestToken, error)
}
