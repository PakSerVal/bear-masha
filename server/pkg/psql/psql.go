package psql

import (
	"crypto/tls"

	"github.com/PakSerVal/bear-masha/internal/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(c config.ConnConfig) (*sqlx.DB, error) {
	pgxConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		return nil, err
	}

	var tlsConfig *tls.Config
	if c.SslMode == "require" {
		tlsConfig = &tls.Config{}
	}
	optConfig := pgx.ConnConfig{
		Host:      c.Host,
		Port:      uint16(c.Port),
		Database:  c.DbName,
		User:      c.User,
		Password:  c.Password,
		TLSConfig: tlsConfig,
	}

	pgxConfig = pgxConfig.Merge(optConfig)

	return sqlx.NewDb(stdlib.OpenDB(pgxConfig), "pgx"), nil
}
