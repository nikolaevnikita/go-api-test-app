package config

import "flag"

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8080
	defaultDbDsn = "postgres://user:password@localhost:5432/go-api-test-app?sslmode=disable"
	defaultMigratePath = "migrations"
)

type Config struct {
	Host  string
	Port  int
	DbDsn string
	MigratePath string
	Debug bool
}

func ReadConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", defaultHost, "flag for explicit server host specifications")
	flag.IntVar(&cfg.Port, "port", defaultPort, "flag for explicit server port specifications")
	flag.StringVar(&cfg.DbDsn, "db", defaultDbDsn, "flag for explicit db connection string")
	flag.StringVar(&cfg.MigratePath, "migrate", defaultMigratePath, "flag for explicit migrate path")
	flag.BoolVar(&cfg.Debug, "debug", false, "flag for explicit debug mode")

	flag.Parse()

	return cfg
}