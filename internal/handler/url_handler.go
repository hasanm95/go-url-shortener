package handler

import "github.com/gin-gonic/gin"

type URLHandler struct{}

func NewURLHandler() *URLHandler{
	return &URLHandler{}
}

func (h *URLHandler) CreateShortURL(c *gin.Context){
	c.JSON(200, gin.H{"message": "Create short URL endpoint"})
}

func (h *URLHandler) RedirectURL(c *gin.Context){
	shortCode := c.Param("shortCode")
	c.JSON(200, gin.H{"message": "Redirect endpoint", "shortCode": shortCode})
}