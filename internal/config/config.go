package config

import (
	"path"
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

func configureViper() (string, error) {
	viper.SetConfigName("stravastats")
	viper.SetConfigType("yaml")

	appPath, err := util.GetApplicationDir()
	if err != nil {
		return "", err
	}

	viper.AddConfigPath(appPath)

	return path.Join(appPath, "stravastats.yaml"), nil
}

func ReadConfig() (Config, error) {
	config := Config{
		Api:     ApiConfig{},
		Display: DisplayConfig{},
	}

	_, err := configureViper()
	if err != nil {
		return config, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config.Api.ClientId = viper.GetString("Api.ClientId")
	config.Api.ClientSecret = viper.GetString("Api.ClientSecret")
	config.Display.Activities = viper.GetStringSlice("Display.Activities")

	return config, nil
}

func SaveConfig(cfg Config) error {
	cfgPath, err := configureViper()
	if err != nil {
		return err
	}

	viper.Set("Api.ClientId", cfg.Api.ClientId)
	viper.Set("Api.ClientSecret", cfg.Api.ClientSecret)

	err = viper.WriteConfigAs(cfgPath)
	if err != nil {
		return err
	}

	return nil
}
