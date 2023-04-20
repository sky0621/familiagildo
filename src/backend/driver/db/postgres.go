package db

import (
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/driver/db/generated"
	"time"
)

type CloseDBClientFunc = func()

func NewQueries(dsn string, option app.DBSetOption) (*generated.Queries, CloseDBClientFunc, error) {
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

	queries := generated.New(db)
	return queries, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Msg("failed to close db")
			}
		}
	}, nil
}
