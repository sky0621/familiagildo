package custommodel

import (
	"io"
)

type Void struct {
}

func (v *Void) MarshalGQL(w io.Writer) {
}

func (v *Void) UnmarshalGQL(a any) error {
	return nil
}
