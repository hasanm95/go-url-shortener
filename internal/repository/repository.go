package repository

import "github.com/hasanm95/go-url-shortener/internal/models"

type URLRepository interface {
	Create(url *models.URL) error
	GetByShortCode(shortCode string) (*models.URL, error)
	IncrementClicks(shortCode string) error
}