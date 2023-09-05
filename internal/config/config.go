package config

import (
	"errors"
	"stravastats/internal/util"

	"github.com/spf13/viper"
)

type Config struct {
	Api     ApiConfig     `mapstructure:"api"`
	Display DisplayConfig `mapstructure:"display"`
}

type ApiConfig struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
}

type DisplayConfig struct {
	Name       string
	Activities []string
}

func configureViper() error {
	viper.SetConfigName("stravastats")
	viper.SetConfigType("yaml")

	appPath, err := util.GetApplicationDir()
	if err != nil {
		return err
	}

	viper.AddConfigPath(appPath)

	return nil
}

func ReadConfig() (*Config, error) {
	err := configureViper()
	if err != nil {
		return nil, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		Api: ApiConfig{
			ClientId:     viper.GetString("Api.ClientId"),
			ClientSecret: viper.GetString("Api.ClientSecret"),
		},
		Display: DisplayConfig{
			Activities: viper.GetStringSlice("Display.Activities"),
		},
	}

	if config.Api.ClientId == "" {
		return nil, errors.New("client id wasn't set")
	}

	if config.Api.ClientSecret == "" {
		return nil, errors.New("client secret wasn't set")
	}

	return config, nil
}

func SaveConfig(cfg *Config) error {
	if cfg == nil {
		return nil
	}

	err := configureViper()
	if err != nil {
		return err
	}

	viper.Set("Api.ClientId", cfg.Api.ClientId)
	viper.Set("Api.ClientSecret", cfg.Api.ClientSecret)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
