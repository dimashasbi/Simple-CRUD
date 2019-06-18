package main

import (
	"M-GateDBConfig/dbsetup"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// migrate DB
	db := dbsetup.DBInit()

	// put configuration A
	dbsetup.DBsetValue(db)
	//
	fmt.Print("Input data is success")
}
