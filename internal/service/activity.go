package service

import (
	"slices"
	"stravastats/internal/api"
	"stravastats/internal/cache"
	"stravastats/internal/model"
	"time"
)

func GetActivities() ([]model.Activity, error) {
	cached := cache.GetActivities()

	from := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	if len(cached) > 0 {
		from = cached[len(cached)-1].StartDate
	}

	activities, err := api.GetActivities(from)
	if err != nil {
		return nil, err
	}

	if len(activities) > 0 {
		cached = append(cached, activities...)
	}

	err = cache.SetActivities(cached)
	if err != nil {
		return nil, err
	}

	return cached, nil
}

func GetActivityTypes() ([]string, error) {
	activities, err := GetActivities()
	if err != nil {
		return nil, err
	}

	var types []string
	for _, a := range activities {
		if !slices.Contains(types, a.Type) {
			types = append(types, a.Type)
		}
	}

	return types, nil
}
