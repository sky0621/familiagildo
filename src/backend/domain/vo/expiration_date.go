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

func ParseExpirationDate(v time.Time) ExpirationDate {
	return ExpirationDate(v)
}
