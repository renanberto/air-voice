package speak

import (
	"github.com/renanberto/air-voice/domain"
	"github.com/renanberto/air-voice/internal"
	"github.com/renanberto/air-voice/model"
)

type Speak struct {
	mysql internal.MysqlRepository
}

func NewSpeaKByID(mysqlRepository internal.MysqlRepository) domain.SpeakUseCase {
	return &Speak{
		mysqlRepository,
	}
}

func (speak *Speak) ByID() {
	var sound model.Sounds
	var sounds []model.Sounds

	query := &model.MysqlQuery{"SELECT * FROM AIR_VOICE.sounds"}
	rows, err := speak.mysql.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&sound.ID, &sound.File, &sound.Location, &sound.Name)
		if err != nil {
			panic(err.Error())
		}
		sounds = append(sounds, sound)
	}
}
