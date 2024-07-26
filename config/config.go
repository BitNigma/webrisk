package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// PG -.
	PG struct {
		URL string `env-required:"true"                 env:"PG_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	log.Println(cfg.Port)

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
