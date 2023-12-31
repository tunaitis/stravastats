package service

import (
	"slices"
	"stravastats/internal/api"
	"stravastats/internal/cache"
	"stravastats/internal/model"
	"strings"
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

	before := len(cached)

	if len(activities) > 0 {
		cached = append(cached, activities...)
	}

	after := len(cached)

	if after > before {
		err = cache.SetActivities(cached)
		if err != nil {
			return nil, err
		}
	}

	return cached, nil
}

func GetActivityTypes(activities []model.Activity) ([]string, error) {
	var types []string
	for _, a := range activities {
		if !slices.Contains(types, a.Type) {
			types = append(types, a.Type)
		}
	}

	return types, nil
}

func GetActivityStats() (*model.Stats, error) {
	activities, err := GetActivities()
	if err != nil {
		return nil, err
	}

	types, err := GetActivityTypes(activities)

	stats := &model.Stats{
		Activities:    make(map[string]model.ActivityStats),
		ActivityTypes: []string{},
		Years:         make(map[int]map[string]model.ActivityStats),
	}

	for _, t := range types {
		key := strings.ToLower(t)

		if !slices.Contains(stats.ActivityTypes, key) {
			stats.ActivityTypes = append(stats.ActivityTypes, key)
		}

		for _, a := range activities {
			if a.Type == t {
				if entry, ok := stats.Activities[key]; ok {
					entry.Distance = entry.Distance + a.Distance
					entry.Duration = entry.Duration + a.Duration
					entry.ElevationGain = entry.ElevationGain + a.ElevationGain
					entry.Count = entry.Count + 1
					stats.Activities[key] = entry
				} else {
					stats.Activities[key] = model.ActivityStats{
						Type:          t,
						Distance:      a.Distance,
						Duration:      a.Duration,
						ElevationGain: a.ElevationGain,
						Count:         1,
					}
				}

				yearKey := a.StartDate.Year()

				if yearEntry, ok := stats.Years[yearKey]; ok {

					if activityEntry, ok := yearEntry[key]; ok {
						activityEntry.Distance = activityEntry.Distance + a.Distance
						activityEntry.Duration = activityEntry.Duration + a.Duration
						activityEntry.ElevationGain = activityEntry.ElevationGain + a.ElevationGain
						activityEntry.Count = activityEntry.Count + 1
						yearEntry[key] = activityEntry
					} else {
						yearEntry[key] = model.ActivityStats{
							Type:          t,
							Distance:      a.Distance,
							Duration:      a.Duration,
							ElevationGain: a.ElevationGain,
							Count:         1,
						}
					}

					stats.Years[yearKey] = yearEntry

				} else {
					stats.Years[yearKey] = make(map[string]model.ActivityStats)
					stats.Years[yearKey][key] = model.ActivityStats{
						Type:          t,
						Distance:      a.Distance,
						Duration:      a.Duration,
						ElevationGain: a.ElevationGain,
						Count:         1,
					}
				}
			}
		}
	}

	return stats, nil
}
