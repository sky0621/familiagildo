package controller

import "github.com/sky0621/familiagildo/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GuildUsecase usecase.GuildInputPort
}

func NewResolver(guild usecase.GuildInputPort) *Resolver {
	return &Resolver{GuildUsecase: guild}
}
