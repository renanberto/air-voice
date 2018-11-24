package domain

import "github.com/renanberto/air-voice/model"

type SpeakUseCase interface {
	ByID(model.Speech) error
}
