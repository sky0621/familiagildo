package mail

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/familiagildo/app"
	"github.com/wneessen/go-mail"
)

type Client struct {
	c         *mail.Client
	from      string
	CloseFunc func()
}

func NewClient(cfg app.Config) (*Client, error) {
	c, err := mail.NewClient(cfg.SmtpHost,
		mail.WithPort(cfg.SmtpPort),
		mail.WithDebugLog(),
		mail.WithTLSPolicy(mail.NoTLS),
		mail.WithTimeout(cfg.SmtpTimeout),
	)
	if err != nil {
		return nil, errors.Join(err)
	}

	closeFunc := func() {
		if c != nil {
			if err := c.Close(); err != nil {
				log.Err(err).Msg("failed to close mail client")
			}
		}
	}

	return &Client{c: c, from: cfg.SmtpFrom, CloseFunc: closeFunc}, nil
}

func (c *Client) GetFrom() string {
	return c.from
}

func (c *Client) SendMail(ctx context.Context, msg *mail.Msg) error {
	return c.c.DialAndSendWithContext(ctx, msg)
}
