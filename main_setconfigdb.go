package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbsetup"
	"fmt"
	"io/ioutil"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	test := true

	// read config file
	test = configuration.getConfig(test)

	// migrate DB
	db := dbsetup.DBInit()

	// put configuration A
	dbsetup.DBsetValue(db)
	configuration.DBsetValue(db)
	//
	// create file
	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Print("Input data is success")
}
