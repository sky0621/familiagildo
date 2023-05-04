package setup

import (
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/mail"
	"github.com/sky0621/familiagildo/driver/web"
)

func NewApp(
	dbClient *db.Client,
	webClient *web.Server,
	mailClient *mail.Client,
) App {
	return App{
		dbClient:   dbClient,
		webClient:  webClient,
		mailClient: mailClient,
	}
}

type App struct {
	dbClient   *db.Client
	webClient  *web.Server
	mailClient *mail.Client
}

func (a *App) Start(webPort string) error {
	if err := a.webClient.S.ListenAndServe(":" + webPort); err != nil {
		return err
	}
	return nil
}

func (a *App) Close() {
	a.webClient.CloseFunc()
	a.dbClient.CloseFunc()
	a.mailClient.CloseFunc()
}
