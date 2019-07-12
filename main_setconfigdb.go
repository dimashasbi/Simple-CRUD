package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbhandler"
	// "M-GateDBConfig/setvalue"
	// "fmt"
	// "io/ioutil"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
type AllModelData struct {
	*configuration.DBConfigurationModel
	*dbsetup.InDB
}

func main() {

	// i want to read file Config here to Connect Database
	myConfig := configuration.GetDBConfig()

	// migrate DB First
	dbsetup.DBInit(&myConfig)

	// open running Server to do Reload Configuration

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
