package cache

import (
	"encoding/json"
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

func getCacheFile() string {
	return path.Join(getUserCacheDir(), "stravastats", "stravastats.json")
}

func RemoveActivities() error {
	file := getCacheFile()

	slog.Debug("Reading cache", slog.String("file", file))

	if _, err := os.Stat(file); os.IsNotExist(err) {
		slog.Debug("Cache not found", slog.String("file", file))
		return nil
	}

	slog.Debug("Removing cache", slog.String("file", file))

	return os.Remove(file)
}

func GetActivities() []model.Activity {
	file := getCacheFile()

	slog.Debug("Reading cache", slog.String("file", file))

	if _, err := os.Stat(file); os.IsNotExist(err) {
		slog.Debug("Cache not found", slog.String("file", file))
		return nil
	}

	reader, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer reader.Close()

	var activities []model.Activity

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&activities)
	if err != nil {
		slog.Debug("An error has occured", slog.String("error", err.Error()))
		return nil
	}

	return activities
}

func SetActivities(activities []model.Activity) error {
	err := createCacheDir()
	if err != nil {
		return err
	}

	file := getCacheFile()

	slog.Debug("Writing cache", slog.String("file", file), slog.Int("count", len(activities)))

	writer, err := os.Create(file)
	if err != nil {
		return err
	}
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	encoder.Encode(activities)

	return nil
}
