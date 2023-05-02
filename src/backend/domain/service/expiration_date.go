package service

import "time"

func CreateGuestTokenExpirationDate() time.Time {
	return time.Now().Add(1 * time.Hour)
}
