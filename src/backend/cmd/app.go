package main

import "github.com/sky0621/kaubandus"

func NewApp() kaubandus.App {
	return &app{}
}

type app struct {
}

func (a *app) Start() error {
	return nil
}
