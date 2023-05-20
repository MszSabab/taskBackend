package config

import "github.com/caarlos0/env/v6"

type Config struct {
	ApplicationPort int    `env:"APPLICATION_PORT"`
	DatabaseUrl     string `env:"DATABASE_URL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
