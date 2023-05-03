package gateway

import (
	"context"
	"fmt"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/event"
	"github.com/sky0621/familiagildo/driver/mail"
	gomail "github.com/wneessen/go-mail"
)

func NewGuildEvent(c *mail.Client) event.GuildEvent {
	return &guildEvent{c: c}
}

type guildEvent struct {
	c *mail.Client
}

func (e *guildEvent) CreateRequested(ctx context.Context, input event.CreateRequestedInput) error {
	m := gomail.NewMsg()
	if err := m.From(e.c.GetFrom()); err != nil {
		return app.WrapError(err, fmt.Sprintf("failed to set From at CreateRequested [token:%s][expirationDate:%v][ownerMail:%s][acceptedNumber:%s]",
			input.Token, input.ExpirationDate, input.OwnerMail, input.AcceptedNumber))
	}
	if err := m.To(input.OwnerMail.ToVal()); err != nil {
		return app.WrapError(err, fmt.Sprintf("failed to set To at CreateRequested [token:%s][expirationDate:%v][ownerMail:%s][acceptedNumber:%s]",
			input.Token, input.ExpirationDate, input.OwnerMail, input.AcceptedNumber))
	}
	// FIXME:
	return nil
}
