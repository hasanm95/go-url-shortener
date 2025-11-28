package service

import (
	"github.com/hasanm95/go-url-shortener/internal/models"
	"github.com/hasanm95/go-url-shortener/internal/repository"
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
	url := &models.URL{
		OriginalURL: originalUrl,
		ShortCode: "abc123",
	}
	s.repository.Create(url)
	return "abc123", nil
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	url, err := s.repository.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}