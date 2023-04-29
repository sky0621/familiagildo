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

func InitializeApp(dsn string, option app.DBSetOption, env app.Env, isTrace bool) (App, error) {
	wire.Build(
		db.NewQueries,

		gateway.NewGuestTokenRepository,
		gateway.NewGuildRepository,
		usecase.NewGuild,

		controller.NewResolver,
		web.NewServer,

		NewApp,
	)
	return App{}, nil
}
