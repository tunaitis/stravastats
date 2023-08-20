package models

import "time"

type Activity struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	Type      string    `json:"type"`
	Distance  float32   `json:"distance"`
	Effort    float32   `json:"suffer_score"`
}
