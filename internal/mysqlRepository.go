package internal

import (
	"github.com/renanberto/air-voice/model"
	"database/sql"
)

type MysqlRepository interface {
	Query(*model.MysqlQuery) (*sql.Rows, error)
}