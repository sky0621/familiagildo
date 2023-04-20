package service

import "github.com/google/uuid"

func CreateToken() string {
	// FIXME:
	return uuid.New().String()
}
