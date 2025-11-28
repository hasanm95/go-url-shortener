package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanm95/go-url-shortener/internal/config"
)

func main(){
	cfg := config.Load()

	r := gin.Default()

  // Define a simple GET endpoint
  r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  port := fmt.Sprintf(":%s", cfg.Port)

  log.Println("Server will run on port %v", port)

  if err := r.Run(port); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}