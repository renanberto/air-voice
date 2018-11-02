package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/renanberto/air-voice/domain"
)

type SpeakHandler struct {
	Speak domain.SpeakUseCase
}

func (this *SpeakHandler) SpeakByID(context *gin.Context) {
	this.Speak.ByID()
}

func NewSpeakHandler(engine *gin.Engine, speakUseCase domain.SpeakUseCase) {
	handler := &SpeakHandler{
		speakUseCase,
	}
	V1 := engine.Group("/v1")
	{
		V1.GET("/speak/:id", handler.SpeakByID)
	}
}
