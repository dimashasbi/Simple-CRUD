package main

import (
	"M-GateDBConfig/dbsetup"
	"fmt"
	"io/ioutil"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// migrate DB
	db := dbsetup.DBInit()

	// put configuration A
	dbsetup.DBsetValue(db)
	//
	// create file
	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Print("Input data is success")
}
