package service

import (
	"crypto/md5"
	"fmt"
	"github.com/sky0621/familiagildo/domain/vo"
)

func CreateAcceptedNumber() vo.AcceptedNumber {
	// FIXME:
	m := md5.New()
	return vo.ParseAcceptedNumber(fmt.Sprintf("%x", m.Sum(nil)))
}
