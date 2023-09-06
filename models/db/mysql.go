package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlProvider struct {
	db *sql.DB
}

func NewMySQLProvider() (SQLProvider, error) {
	banco := MysqlProvider{}
	erro := banco.Connect()

	return &banco, erro
}

func (m *MysqlProvider) Connect() error {
	conectionString := "root:root@tcp(localhost:3307)/"
	db, err := sql.Open("mysql", conectionString)
	if err != nil {
		fmt.Sprintf("Erro ao conectar com o banco! %v", err)
	}

	defer db.Close()

	return db.Ping()
}

type SQLProvider interface {
	Connect() error
}
