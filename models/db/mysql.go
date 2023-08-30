package db

import (
	"database/sql"
	"fmt"
)

type MysqlProvider struct {
	db *sql.DB
}

func NewMySQLProvider() (SQLProvider, error) {
	provider := MysqlProvider{}

	return &provider, provider.Connect()
}

func (m *MysqlProvider) Connect() error {
	conectionString := "root:root@tcp(localhost:3308)?charset=utf8&parseTime=True&loc=Local"
	_, err := sql.Open("mysql", conectionString)
	if err != nil {
		fmt.Errorf("Erro ao conectar com o banco! %v", err)
	}
	return m.db.Ping()
}

type SQLProvider interface {
	Connect() error
}
