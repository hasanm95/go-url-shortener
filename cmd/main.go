package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanm95/go-url-shortener/internal/config"
	"github.com/hasanm95/go-url-shortener/internal/database"
	"github.com/hasanm95/go-url-shortener/internal/handler"
	"github.com/hasanm95/go-url-shortener/internal/repository"
	"github.com/hasanm95/go-url-shortener/internal/router"
	"github.com/hasanm95/go-url-shortener/internal/service"
)

func main(){
	cfg := config.Load()

  db, err := database.NewPostgresDB(cfg.DatabaseURL)

  if err != nil {
    log.Println(err)
    return;
  }

  repo := repository.NewPostgresRepository(db)

  urlService := service.NewURLService(repo)
  urlHandler := handler.NewURLHandler(urlService)

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