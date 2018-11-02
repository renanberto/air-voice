package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renanberto/air-voice/internal"
	"github.com/renanberto/air-voice/model"
	"os"
)

type MysqlRepository struct {
	conn *sql.DB
}

func NewMysqlRepository() internal.MysqlRepository {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	conn, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/")
	if err != nil {
		panic(err)
	}
	return &MysqlRepository{conn}
}

func (m *MysqlRepository) Query(mysql *model.MysqlQuery) (*sql.Rows, error) {
	return m.conn.Query(mysql.Query)
}
