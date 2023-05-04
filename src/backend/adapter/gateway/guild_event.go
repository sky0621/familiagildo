package gateway

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/event"
	"github.com/sky0621/familiagildo/driver/mail"
	gomail "github.com/wneessen/go-mail"
	"text/template"
)

//go:embed template/requestCreateGuildByGuest.tmpl
var requestCreateGuildByGuestTmpl string

const (
	requestCreateGuildByGuestMailTitle = "ギルド仮登録完了メール"
	requestCreateGuildByGuestTmplName  = "requestCreateGuildByGuest.tmpl"
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
		return app.WrapError(err, fmt.Sprintf("failed to set From at CreateRequested [token:%s][from:%s]",
			input.Token, e.c.GetFrom()))
	}

	if err := m.To(input.OwnerMail.ToVal()); err != nil {
		return app.WrapError(err, fmt.Sprintf("failed to set To at CreateRequested [token:%s][to:%s]",
			input.Token, input.OwnerMail.ToVal()))
	}

	m.Subject(requestCreateGuildByGuestMailTitle)

	t := template.Must(template.New("requestCreateGuildByGuest.tmpl").Parse(requestCreateGuildByGuestTmpl))
	if err := m.SetBodyTextTemplate(t, input); err != nil {
		return app.WrapError(err, fmt.Sprintf("failed to set Body at CreateRequested [input: %s]", input))
	}

	if err := e.c.SendMail(ctx, m); err != nil {
		return app.WrapError(err, fmt.Sprintf("failed to send Mail at CreateRequested [input: %s]", input))
	}

	return nil
}
