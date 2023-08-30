package main

import (
	"fmt"
	"um_por_cento/helper"
	"um_por_cento/models/db"
)

func main() {
	helper.Help()
	mysqlprovider, err := db.NewMySQLProvider()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mysqlprovider)
}
