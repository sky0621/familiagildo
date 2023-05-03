//go:build wireinject

package setup

import (
	"github.com/google/wire"
	"github.com/sky0621/familiagildo/adapter/controller"
	"github.com/sky0621/familiagildo/adapter/gateway"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/mail"
	"github.com/sky0621/familiagildo/driver/web"
	"github.com/sky0621/familiagildo/usecase"
)

func InitializeApp(cfg app.Config) (App, error) {
	wire.Build(
		db.NewQueries,
		mail.NewClient,

		gateway.NewTransactionRepository,
		gateway.NewGuestTokenRepository,
		gateway.NewGuildRepository,
		gateway.NewGuildEvent,
		usecase.NewGuild,

		controller.NewResolver,
		web.NewServer,

		NewApp,
	)
	return App{}, nil
}
