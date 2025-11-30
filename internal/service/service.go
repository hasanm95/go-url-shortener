package service

import (
	"github.com/hasanm95/go-url-shortener/internal/models"
	"github.com/hasanm95/go-url-shortener/internal/repository"
	"github.com/hasanm95/go-url-shortener/internal/utils"
)

type URLService struct {
	repository repository.URLRepository
}

func NewURLService(repository repository.URLRepository) *URLService{
	return &URLService{
		repository: repository,
	}
}

func (s *URLService) CreateShortURL(originalUrl string) (string, error){
	// Step 1: Create URL with temp short code
	url := &models.URL{
		OriginalURL: originalUrl,
		ShortCode: "temp",
	}

	// Step 2: Insert into DB to get auto-generated ID
	err := s.repository.Create(url)
	if err != nil {
		return "", err
	}

	// Step 3: Generate base62 short code from ID
	url.ShortCode = utils.EncodeToBase62(url.ID)

	// Step 4: Update the short code in the DB
	s.repository.UpdateShortCode(url.ID, url.ShortCode)
	return url.ShortCode, nil
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	url, err := s.repository.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}

	// Increment clicks
	s.repository.IncrementClicks(shortCode)

	return url.OriginalURL, nil
}

func (s *URLService) RetriveOriginalURL(shortCode string) (*models.URL, error) {
	url, err := s.repository.GetByShortCode(shortCode)

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) UpdateShortURL(shortCode string, url string)(*models.URL, error){
	urlObj, err := s.repository.UpdateShortURL(shortCode, url)

	if err != nil {
		return nil, err
	}

	return urlObj, nil
}


func (s *URLService) DeleteShortURL(shortCode string) error {
	err := s.repository.DeleteShortURL(shortCode)

	if err != nil {
		return err
	}

	return nil
}