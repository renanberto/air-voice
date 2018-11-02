package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/renanberto/air-voice/domain/speak"
	"github.com/renanberto/air-voice/internal/handlers"
	"github.com/renanberto/air-voice/internal/repository"
	"net/http"
	"time"
)

func NewHttpHandler() *gin.Engine {
	httpHandler := gin.Default()
	httpHandler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message:": "OK"})
	})
	httpHandler.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	return httpHandler
}

func main() {
	handler := NewHttpHandler()

	// Repositories
	mysqlRepository := repository.NewMysqlRepository()

	// UseCases
	speakUseCase := speak.NewSpeaKByID(mysqlRepository)

	handlers.NewSpeakHandler(handler, speakUseCase)
	handler.Run(":4000")
}
