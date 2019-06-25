package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbsetup"
	"M-GateDBConfig/setvalue"
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

	// migrate DB
	db := dbsetup.DBInit(&myConfig)

	// put configuration A
	setvalue.SetValue(db)

	// create file
	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	defer db.Close()
	fmt.Print("Input data is success")
}
