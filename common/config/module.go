package config

import (
	cfg "auth/config"
	"go.uber.org/fx"
)

func NewConfig() cfg.Config {
	return cfg.NewViper(
		cfg.WithRequiredConfig([]string{
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASS",
		}),
	)
}

var Module = fx.Provide(NewConfig)
