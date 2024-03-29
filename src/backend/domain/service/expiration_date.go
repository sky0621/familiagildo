package service

import (
	"github.com/sky0621/familiagildo/domain/vo"
	"time"
)

func CreateGuestTokenExpirationDate() vo.ExpirationDate {
	return vo.ToExpirationDate(time.Now().Add(24 * time.Hour))
}
