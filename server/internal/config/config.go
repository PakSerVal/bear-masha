package config

import (
	"flag"
)

type Config struct {
	HTTPPort  int64
	IsDev     bool
	Password  string
	Login     string
	SecretKey string
	DbConn    ConnConfig
}

type ConnConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	DbName   string
	SslMode  string
}

func New() *Config {
	c := &Config{}

	flag.Int64Var(&c.HTTPPort, "httpPort", 8080, "Http port")
	flag.BoolVar(&c.IsDev, "dev", true, "development mode")
	flag.StringVar(&c.SecretKey, "secretKey", "", "secret key")

	flag.StringVar(&c.Login, "login", "admin", "login")
	flag.StringVar(&c.Password, "password", "admin", "password")
	flag.StringVar(&c.DbConn.Host, "pgHost", "localhost", "postgres host")
	flag.Int64Var(&c.DbConn.Port, "pgPort", 5432, "postgres port")
	flag.StringVar(&c.DbConn.User, "pgUser", "postgres", "postgres user")
	flag.StringVar(&c.DbConn.Password, "pgPass", "postgres", "postgres password")
	flag.StringVar(&c.DbConn.DbName, "pgDatabase", "postgres", "postgres database name")
	flag.StringVar(&c.DbConn.SslMode, "pgSslMode", "disable", "postgres ssl mode")

	flag.Parse()

	return c
}
