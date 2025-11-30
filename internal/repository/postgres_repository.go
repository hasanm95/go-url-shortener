package repository

import (
	"database/sql"
	"fmt"

	"github.com/hasanm95/go-url-shortener/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository  {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(url *models.URL) error {
	query := `INSERT INTO urls (original_url, short_code) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(query, url.OriginalURL, url.ShortCode).Scan(&url.ID)

	if err != nil {
		return fmt.Errorf("error creating url: %w", err)
	}

	return nil
}

func (r *PostgresRepository) GetByShortCode(shortCode string) (*models.URL, error) {
	url := &models.URL{}
	query := `SELECT id, original_url, short_code, created_at, clicks FROM urls WHERE short_code = $1`

	err := r.db.QueryRow(query, shortCode).Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.CreatedAt, &url.Clicks)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("url not found")
		}
		return nil, fmt.Errorf("error getting url: %w", err)
	}

	return url, nil
}

func (r *PostgresRepository) IncrementClicks(shortCode string) error {
	query := `UPDATE urls SET clicks = clicks + 1 WHERE short_code = $1`
	_, err := r.db.Exec(query, shortCode)
	if err != nil {
		return fmt.Errorf("error incrementing clicks: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateShortCode(id int, shortCode string) error {
	query := `UPDATE urls SET short_code = $2 WHERE id = $1`
	_, err := r.db.Exec(query, id, shortCode)

	if err != nil {
		return fmt.Errorf("error updating short code: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateShortURL(shortCode string, url string) (*models.URL, error) {
	query := `UPDATE urls SET original_url = $2 WHERE id = $1`
	_, err := r.db.Exec(query, shortCode, url)

	if err != nil {
		return nil, fmt.Errorf("error updating short code: %w", err)
	}
	urlObj, err := r.GetByShortCode(shortCode)

	if err != nil {
		return nil, fmt.Errorf("error fetching short code: %w", err)
	}
	return urlObj, nil
}