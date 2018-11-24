package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/renanberto/air-voice/domain"
	"github.com/renanberto/air-voice/model"
	"net/http"
)

type SpeakHandler struct {
	Speak domain.SpeakUseCase
}

func (this *SpeakHandler) SpeakByID(context *gin.Context) {
	speech := model.Speech{}
	err := context.BindJSON(&speech)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"error": "Cannot bind properts",
		})
		context.Abort()
		return
	}
	speech.ID = context.Param("id")
	speech.Organization = context.Param("organization")
	this.Speak.ByID(speech)
}

func NewSpeakHandler(engine *gin.Engine, speakUseCase domain.SpeakUseCase) {
	handler := &SpeakHandler{
		speakUseCase,
	}
	V1 := engine.Group("/v1")
	{
		V1.POST("/speak/:organization/:id", handler.SpeakByID)
	}
}
