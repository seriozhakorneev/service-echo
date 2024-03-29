package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		Rewriter `yaml:"rewriter"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Rewriter -.
	Rewriter struct {
		Active bool  `yaml:"active"   env:"ACTIVE"`
		Rules  Rules `yaml:"rules"`
	}

	// Rules -.
	Rules struct {
		Name  string `env-required:"true" yaml:"name"   env:"NAME"`
		Value string `env-required:"true" yaml:"value"   env:"VALUE"`
		New   string `env-required:"true" yaml:"new"   env:"NEW"`
	}
)

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
