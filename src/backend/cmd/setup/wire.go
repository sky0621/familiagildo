package setup

import (
	"github.com/google/wire"
	"github.com/sky0621/familiagildo/adapter/gateway"
	"github.com/sky0621/familiagildo/usecase"
)

func InitializeApp() {
	wire.Build(
		gateway.NewGuildRepository,
		gateway.NewTaskRepository,
		usecase.NewGuild,
	)
}
