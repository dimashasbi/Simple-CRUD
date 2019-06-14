package main

import (
	"M-Gate/setconfigdb/dbsetup"

	"github.com/jinzhu/gorm"
)

func main() {
	// migrate DB
	db := dbsetup.DBInit()

	// put configuration A
	dbsetup.DBsetValue()
	//
}
