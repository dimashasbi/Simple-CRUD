package main

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/handler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
type AllModelData struct {
	dbConfig   *configuration.DBConfigurationModel
	DBHandler  *dbhandler.DBHandler
	AppHandler *handler.AppHandler
}

func main() {

	myCoreVariable := &AllModelData{}
	// i want to read file Config here to Connect Database
	dbConfig, err := configuration.GetDBConfig()
	checkErr(err)
	// migrate DB First
	db := dbhandler.DBInit(dbConfig)

	// open running Server to do Reload Configuration
	handler := &handler.AppHandler{}
	handler.Initialize(db)
	handler.Run(":4911")

	myCoreVariable.dbConfig = &dbConfig
	myCoreVariable.DBHandler = db
	myCoreVariable.AppHandler = handler
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
