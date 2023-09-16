package config

import (
	"fmt"
	"path"
	"reflect"
	"stravastats/internal/util"
	"strings"

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

func SetValue(config *Config, name string, value string) error {
	if name == "Api.ClientId" {
		config.Api.ClientId = value
		return nil
	}

	if name == "Api.ClientSecret" {
		config.Api.ClientSecret = value
		return nil
	}

	if name == "Display.Activities" {
		config.Display.Activities = strings.Split(value, ",")
		return nil
	}

	return fmt.Errorf("variable not found: %s", name)
}

func GetValue(data interface{}, index string) (string, error) {
	indexArray := strings.Split(index, ".")

	r := reflect.ValueOf(data)
	for _, i := range indexArray {
		if r.FieldByName(i).Kind() == reflect.Struct {
			r = reflect.ValueOf(r.FieldByName(i).Interface())
		} else {
			r = r.FieldByName(i)
		}
	}

	if r.IsValid() {
		return r.String(), nil
	}

	return "", fmt.Errorf("variable not found: %s", index)
}

func (c *ApiConfig) SetValue(name string, value string) {

}

func GetConfigPath() (string, error) {
	appPath, err := util.GetApplicationDir()
	if err != nil {
		return "", err
	}

	return path.Join(appPath, "stravastats.yaml"), nil
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
	viper.Set("Display.Activities", cfg.Display.Activities)

	err = viper.WriteConfigAs(cfgPath)
	if err != nil {
		return err
	}

	return nil
}
