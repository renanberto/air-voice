package internal

import (
	"database/sql"
)

type MysqlRepository interface {
	Query(string) (*sql.Rows, error)
}