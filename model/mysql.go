package model

type MysqlQuery struct {
	Query  string
}

type Sounds struct {
	ID       int
	Name     string
	Location string
	File     string
}
