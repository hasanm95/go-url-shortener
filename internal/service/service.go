package service

type URLService struct {}

func NewURLService() *URLService{
	return &URLService{}
}

func (s *URLService) CreateShortURL(originalUrl string) (string, error){
	return "abc123", nil
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	return "https://google.com", nil
}