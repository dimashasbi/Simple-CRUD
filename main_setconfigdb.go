package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/handler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
type AllModelData struct {
	dbConfig    *configuration.DBConfigurationModel
	DBHandler   *dbhandler.DBHandler
	HTTPHandler *handler.HTTPHandler
}

func main() {

	myCoreVariable := &AllModelData{}
	// i want to read file Config here to Connect Database
	dbConfig := configuration.GetDBConfig()

	// migrate DB First
	db := dbhandler.DBInit(dbConfig)

	// open running Server to do Reload Configuration
	handler := &handler.HTTPHandler{}
	handler.Initialize(db)
	handler.Run(":4911")

	myCoreVariable.dbConfig = &dbConfig
	myCoreVariable.DBHandler = db
	myCoreVariable.HTTPHandler = handler
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
