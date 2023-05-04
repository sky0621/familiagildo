package service

import (
	"fmt"
	"github.com/sky0621/familiagildo/domain/vo"
	"math/rand"
	"time"
)

func CreateAcceptedNumber() vo.AcceptedNumber {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return vo.ParseAcceptedNumber(fmt.Sprintf("%010d", r.Int31()))
}
