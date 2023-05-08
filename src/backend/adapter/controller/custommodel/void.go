package custommodel

import (
	"fmt"
	"io"
)

type Void struct {
}

func (v *Void) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, "{}")
	if err != nil {
		fmt.Println(err)
	}
}

func (v *Void) UnmarshalGQL(a any) error {
	fmt.Println("UnmarshalGQL")
	return nil
}
