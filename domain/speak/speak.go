package speak

import (
	"github.com/renanberto/air-voice/domain"
	"github.com/renanberto/air-voice/internal"
	"github.com/renanberto/air-voice/model"
)

type Speak struct {
	mysql internal.MysqlRepository
	speech internal.SpeechRepository
}

func NewSpeaKByID(mysqlRepository internal.MysqlRepository, speechRepository internal.SpeechRepository) domain.SpeakUseCase {
	return &Speak{
		mysqlRepository,
		speechRepository,

	}
}

func (speak *Speak) ByID(speech model.Speech) error {

	speak.mysql.Query("")

	sound, err := speak.speech.TTS(speech.Content)
	if err != nil {
		return err
	}

	err = speak.speech.PlaySound(sound)
	return err
}