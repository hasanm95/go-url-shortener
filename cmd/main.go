package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanm95/go-url-shortener/internal/config"
	"github.com/hasanm95/go-url-shortener/internal/handler"
	"github.com/hasanm95/go-url-shortener/internal/router"
)

func main(){
	cfg := config.Load()

  urlHandler := handler.NewURLHandler()

	r := router.SetupRouter(urlHandler)


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pong",
		})
	})

  port := fmt.Sprintf(":%s", cfg.Port)

  log.Printf("Server will run on port %v", port)

  if err := r.Run(port); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}