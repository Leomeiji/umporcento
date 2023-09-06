package main

import (
	"fmt"

	"um_por_cento/models/db"
)

func main() {
	_, err := db.NewMySQLProvider()
	if err != nil {
		panic(err)
	}

	fmt.Println("conectou com o banco")
}
