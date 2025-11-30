package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanm95/go-url-shortener/internal/handler"
)


func SetupRouter(urlHandler *handler.URLHandler) *gin.Engine{
	r := gin.Default()

	r.POST("/shorten", urlHandler.CreateShortURL)
	r.GET("/shorten/:shortCode", urlHandler.RetriveOriginalURL)
	r.PUT("/shorten/:shortCode", urlHandler.UpdateShortURL)
	r.GET("/:shortCode", urlHandler.RedirectURL)

	return r
}