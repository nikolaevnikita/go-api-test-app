package config

import "flag"

type Config struct {
	Host  string
	Port  int
	Debug bool
}

func ReadConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", "localhost", "flag for explicit server host specifications")
	flag.IntVar(&cfg.Port, "port", 8080, "flag for explicit server port specifications")
	flag.BoolVar(&cfg.Debug, "debug", false, "flag for explicit debug mode")

	flag.Parse()

	return cfg
}