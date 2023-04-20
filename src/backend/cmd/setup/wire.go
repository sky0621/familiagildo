//go:build wireinject

package setup

import (
	"github.com/google/wire"
	"github.com/sky0621/familiagildo/adapter/controller"
	"github.com/sky0621/familiagildo/adapter/gateway"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/web"
	"github.com/sky0621/familiagildo/usecase"
)

func InitializeQueries()
func InitializeApp(cfg app.Config) (App, error) {
	wire.Build(
		app.ToDsn,
		app.ToDBSetOption,
		app.GetEnv,
		app.IsTrace,
		db.NewQueries,

		gateway.NewGuildRepository,
		gateway.NewTaskRepository,
		usecase.NewGuild,

		controller.NewResolver,
		web.NewServer,
	)
	return App{}, nil
}
