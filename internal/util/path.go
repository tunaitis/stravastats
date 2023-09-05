package util

import (
	"os"
	"path/filepath"
)

func GetApplicationDir() (string, error) {
	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	app := filepath.Join(cfg, ".stravastats")
	if _, err := os.Stat(app); os.IsNotExist(err) {
		err := os.MkdirAll(app, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return app, nil
}
