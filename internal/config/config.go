package config

import (
	"cmp"
	"flag"
	"os"
	"strconv"

	"github.com/nikolaevnikita/go-api-test-app/internal/logger"
)

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

func ReadConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", defaultHost, "flag for explicit server host specifications")
	flag.IntVar(&cfg.Port, "port", defaultPort, "flag for explicit server port specifications")
	flag.StringVar(&cfg.DbDsn, "db", defaultDbDsn, "flag for explicit db connection string")
	flag.StringVar(&cfg.MigratePath, "migrate", defaultMigratePath, "flag for explicit migrate path")
	flag.BoolVar(&cfg.Debug, "debug", false, "flag for explicit debug mode")

	flag.Parse()

	if cfg.Host == defaultHost {
		cfg.Host = cmp.Or(os.Getenv("HOST"), cfg.Host)
	}
	if cfg.Port == defaultPort {
		envPortValue := os.Getenv("PORT")
		envPort, err := strconv.Atoi(envPortValue)
		if err != nil {
			if envPortValue != "" {
				log := logger.Get()
				log.Debug().Err(err).Str("PORT", envPortValue).Msg("incorrect PORT environment variable")
			}
		} else {
			cfg.Port = envPort
		}
	}
	if cfg.DbDsn == defaultDbDsn {
		cfg.DbDsn = cmp.Or(os.Getenv("DB_DSN"), cfg.DbDsn)
	}
	if cfg.MigratePath == defaultMigratePath {
		cfg.MigratePath = cmp.Or(os.Getenv("MIGRATE_PATH"), cfg.MigratePath)
	}

	return &cfg
}