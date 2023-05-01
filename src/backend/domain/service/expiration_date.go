package service

import "time"

func CreateExpirationDate() time.Time {
	// FIXME:
	return time.Now().Add(7 * 24 * time.Hour)
}
