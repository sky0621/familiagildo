package postgres

import (
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/sky0621/kaubandus/adapter/gateway/ent"
	"github.com/sky0621/kaubandus/cmd/setup"
	"time"
)

type CloseDBClientFunc = func()

func Open(dsn string, option setup.DBSetOption) (*ent.Client, CloseDBClientFunc, error) {
	var connector driver.Connector
	connector, err := pq.NewConnector(dsn)
	if err != nil {
		return nil, func() {}, errors.Join(err)
	}
	connector = ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())

	db := sql.OpenDB(connector)
	db.SetMaxIdleConns(option.DBMaxIdleConnections)
	db.SetMaxOpenConns(option.DBMaxOpenConnections)
	db.SetConnMaxLifetime(time.Duration(option.DBConnMaxLifetimeMinutes) * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, nil, errors.Join(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), func() {
		if drv != nil {
			drv.Close()
		}
	}, nil
}
