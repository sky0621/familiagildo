package main

import (
	"github.com/google/wire"
	"gocloud.dev/server"
)

func init() {
	_ = appSet
}

var appSet = wire.NewSet(
	newApp,
)

type app struct {
	s *server.Server
}

func newApp(s *server.Server) *app {
	return &app{s: s}
}
