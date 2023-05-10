package entity

import "github.com/sky0621/familiagildo/domain/vo"

type GuestToken struct {
	ID             vo.ID
	Token          vo.Token
	ExpirationDate vo.ExpirationDate
	Mail           vo.OwnerMail
}
