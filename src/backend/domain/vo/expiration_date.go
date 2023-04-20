package vo

import "time"

type ExpirationDate time.Time

func (v ExpirationDate) Validate() bool {
	// FIXME:
	return true
}

func (v ExpirationDate) ToVal() time.Time {
	return time.Time(v)
}

func ParseExpirationDate(v time.Time) ExpirationDate {
	return ExpirationDate(v)
}
