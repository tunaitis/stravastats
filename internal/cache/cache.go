package cache

import (
	"log/slog"
	"os"
	"path"
	"stravastats/internal/model"

	"fmt"
	"runtime"
)

func getUserCacheDir() string {
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s/Library/Caches", os.Getenv("HOME"))
	case "linux":
		return fmt.Sprintf("%s/.cache", os.Getenv("HOME"))
	case "windows":
		return fmt.Sprintf("C:\\Users\\%%%s%%\\AppData\\Local", os.Getenv("USER"))
	}
	return ""
}

func createCacheDir() error {
	cachePath := path.Join(getUserCacheDir(), "stravastats")
	return os.MkdirAll(cachePath, 0700)
}

func GetActivities() []model.Activity {
	file := path.Join(getUserCacheDir(), "stravastats", "stravastats.json")

	slog.Debug("Reading cache", slog.String("file", file))

	if _, err := os.Stat(file); os.IsNotExist(err) {
		slog.Debug("Cache not found", slog.String("file", file))
		return nil
	}

	return nil
}

func SetActivities(activities []model.Activity) error {
	err := createCacheDir()
	if err != nil {
		return err
	}

	return nil
}
