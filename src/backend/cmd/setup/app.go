package setup

import (
	"errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/web"
	"gocloud.dev/server"
)

func NewApp(
	cfg app.Config,
	closeDBClientFunc db.CloseDBClientFunc,
) (*App, error) {
	queries, closeDBClientFunc, err := db.NewQueries(cfg.Dsn(), cfg.ToDBSetOption())
	if err != nil {
		return nil, errors.Join(err)
	}
	return &App{
		Server:            server,
		CloseDBClientFunc: closeDBClientFunc,
		CloseServerFunc:   closeServerFunc,
	}
}

type App struct {
	Server            *server.Server
	CloseDBClientFunc db.CloseDBClientFunc
	CloseServerFunc   web.CloseServerFunc
}

func (a *App) Start(webPort string) error {
	if err := a.Server.ListenAndServe(":" + webPort); err != nil {
		return err
	}
	return nil
}

func (a *App) Close() {
	a.CloseDBClientFunc()
	a.CloseServerFunc()
}
