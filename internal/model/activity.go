package model

import "time"

type Activity struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	Type      string    `json:"type"`
	Distance  float32   `json:"distance"`
	Effort    float32   `json:"suffer_score"`
}

type ActivityStats struct {
	Type     string
	Distance float32
}

type Stats struct {
	Activities map[string]ActivityStats
	Years      map[int]map[string]ActivityStats
}
