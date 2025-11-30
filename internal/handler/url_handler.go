package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanm95/go-url-shortener/internal/service"
)

type URLHandler struct{
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler{
	return &URLHandler{
		service: service,
	}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func (h *URLHandler) CreateShortURL(c *gin.Context){
	var req ShortenRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	shortCode, err := h.service.CreateShortURL(req.URL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_code": shortCode,
		"short_url":  "http://localhost:8080/" + shortCode,
	})

}

func (h *URLHandler) RedirectURL(c *gin.Context){
	shortCode := c.Param("shortCode")
	
	originalUrl, err := h.service.GetOriginalURL(shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
	}

	c.Redirect(http.StatusMovedPermanently, originalUrl)
}

func (h *URLHandler) RetriveOriginalURL(c *gin.Context){
	shortCode := c.Param("shortCode")
	url, err := h.service.RetriveOriginalURL(shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
	}

	c.JSON(http.StatusOK, url)
}