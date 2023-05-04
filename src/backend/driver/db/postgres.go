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

type Client struct {
	Q         *generated.Queries
	CloseFunc func()
	DB        *sql.DB
}

func NewClient(cfg app.Config) (*Client, error) {
	var connector driver.Connector
	connector, err := pq.NewConnector(cfg.Dsn())
	if err != nil {
		return nil, errors.Join(err)
	}
	connector = ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())

	option := cfg.ToDBSetOption()

	db := sql.OpenDB(connector)
	db.SetMaxIdleConns(option.DBMaxIdleConnections)
	db.SetMaxOpenConns(option.DBMaxOpenConnections)
	db.SetConnMaxLifetime(time.Duration(option.DBConnMaxLifetimeMinutes) * time.Minute)
	if err := db.Ping(); err != nil {
		return nil, errors.Join(err)
	}

	queries := generated.New(db)

	closeFunc := func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Msg("failed to close db")
			}
		}
	}

	return &Client{Q: queries, CloseFunc: closeFunc, DB: db}, nil
}
