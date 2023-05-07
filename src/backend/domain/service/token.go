package service

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/sky0621/familiagildo/domain/vo"
	"io"
)

func CreateToken() vo.Token {
	b := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return vo.ToToken(hex.EncodeToString(b))
}
