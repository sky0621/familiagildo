package controller

import "github.com/sky0621/familiagildo/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Guild usecase.Guild
}

func NewResolver(guild usecase.Guild) *Resolver {
	return &Resolver{Guild: guild}
}
