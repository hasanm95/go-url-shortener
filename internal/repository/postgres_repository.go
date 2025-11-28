package repository

import (
	"database/sql"

	"github.com/hasanm95/go-url-shortener/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository  {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(url *models.URL) error {
	return nil
}

func (r *PostgresRepository) GetByShortCode(shortCode string) (*models.URL, error) {
	return &models.URL{OriginalURL: "https://google.com", ShortCode: "abc123"}, nil
}

func (r *PostgresRepository) IncrementClicks(shortCode string) error {
	return nil
}