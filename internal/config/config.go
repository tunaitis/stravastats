package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Api ApiConfig `mapstructure:"api"`
}

type ApiConfig struct {
	ClientId     string
	ClientSecret string
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("stravastats")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.stravastats")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if config.Api.ClientId == "" {
		return nil, errors.New("client id wasn't set")
	}

	if config.Api.ClientSecret == "" {
		return nil, errors.New("client secret wasn't set")
	}

	return nil, nil
}
