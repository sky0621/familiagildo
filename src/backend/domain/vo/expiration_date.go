package vo

import (
	"time"
)

type ExpirationDate time.Time

func (v ExpirationDate) Validate() error {
	// FIXME:
	return nil
}

func (v ExpirationDate) FieldName() string {
	return "expirationDate"
}

func (v ExpirationDate) ToVal() time.Time {
	return time.Time(v)
}

func ToExpirationDate(v time.Time) ExpirationDate {
	return ExpirationDate(v)
}

func ParseExpirationDate(v time.Time) (ExpirationDate, error) {
	ed := ToExpirationDate(v)
	if err := ed.Validate(); err != nil {
		return ed, err
	}
	return ed, nil
}
