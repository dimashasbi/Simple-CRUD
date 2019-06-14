package main

import (
	"M-GateDBConfig/dbsetup"

	"github.com/jinzhu/gorm"
)

func main() {
	// migrate DB
	db := dbsetup.DBInit()

	// put configuration A
	dbsetup.DBsetValue()
	//
}
