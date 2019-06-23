package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbsetup"
	"fmt"
	"io/ioutil"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
// type AllModelData struct {
// 	*configuration.ConfigurationModel
// 	*dbsetup.InDB
// }

func main() {

	// i want to read file Config here
	myConfig := configuration.GetConfig()
	fmt.Printf("isinya : %v \n", myConfig.Password)
	// migrate DB
	db := dbsetup.DBInit(&myConfig)

	// put configuration A
	dbsetup.DBsetValue(db)
	// configuration.DBsetValue(db)
	//
	// create file
	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Print("Input data is success")
}
