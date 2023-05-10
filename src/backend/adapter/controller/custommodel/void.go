package custommodel

import (
	"github.com/sky0621/familiagildo/app/log"
	"io"
)

type Void struct {
}

func (v *Void) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, "{}")
	if err != nil {
		log.ErrorSend(err)
	}
}

func (v *Void) UnmarshalGQL(a any) error {
	log.Infof("UnmarshalGQL: %#v", a)
	return nil
}
