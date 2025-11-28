package models

import "time"

type URL struct {
	ID int `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode string 	`json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	Clicks int `json:"clicks"`
}