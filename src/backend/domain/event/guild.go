package event

import (
	"context"
	"fmt"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildEvent interface {
	CreateRequested(ctx context.Context, input CreateRequestedInput) error
}

type CreateRequestedInput struct {
	GuildName      vo.GuildName
	Token          vo.Token
	ExpirationDate vo.ExpirationDate
	OwnerMail      vo.OwnerMail
	AcceptedNumber vo.AcceptedNumber
}

func (i CreateRequestedInput) String() string {
	return fmt.Sprintf("CreateRequestedInput { GuildName:%s, Token:%s, ExpirationDate:%v, OwnerMail:%s, AcceptedNumber:%s }",
		i.GuildName.ToVal(), i.Token.ToVal(), i.ExpirationDate.ToVal(), i.OwnerMail.ToVal(), i.AcceptedNumber.ToVal())
}
