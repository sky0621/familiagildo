package mail

import (
	"errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/wneessen/go-mail"
)

func NewClient(cfg app.Config) (*Client, error) {
	c, err := mail.NewClient(cfg.SmtpHost, mail.WithPort(cfg.SmtpPort))
	if err != nil {
		return nil, errors.Join(err)
	}
	return &Client{c: c, from: cfg.SmtpFrom}, nil
}

type Client struct {
	c    *mail.Client
	from string
}

func (c *Client) GetFrom() string {
	return c.from
}
