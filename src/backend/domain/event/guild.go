package event

import (
	"context"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildEvent interface {
	CreateRequested(ctx context.Context, input CreateRequestedInput) error
}

type CreateRequestedInput struct {
	Token          vo.Token
	ExpirationDate vo.ExpirationDate
	OwnerMail      vo.OwnerMail
	AcceptedNumber vo.AcceptedNumber
}
