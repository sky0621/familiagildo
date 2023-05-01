package service

import (
	"crypto/md5"
	"fmt"
)

func CreateAcceptedNumber() string {
	// FIXME:
	m := md5.New()
	return fmt.Sprintf("%x", m.Sum(nil))
}
