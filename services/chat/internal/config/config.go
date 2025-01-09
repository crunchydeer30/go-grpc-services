package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	GRPC     GRPCConfig     `validate:"required"`
	Postgres PostgresConfig `validate:"required"`
}

type GRPCConfig struct {
	Host string `validate:"required"`
	Port string `validate:"required"`
}

type PostgresConfig struct {
	DSN string
}

func Load(configPath string) (*Config, error) {
	v := viper.New()

	var cfg Config
	v.SetConfigFile(configPath)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(cfg)
}
