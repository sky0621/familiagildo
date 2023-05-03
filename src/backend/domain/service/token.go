package service

import (
	"github.com/google/uuid"
	"github.com/sky0621/familiagildo/domain/vo"
)

func CreateToken() vo.Token {
	// FIXME:
	return vo.ParseToken(uuid.New().String())
}
