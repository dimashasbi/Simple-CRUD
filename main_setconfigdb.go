package main

import (
	"M-GateDBConfig/adapter"
	"M-GateDBConfig/app/fileconfig"
	"M-GateDBConfig/model"

	"fmt"

	"M-GateDBConfig/provider/dbhandler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
type AllModelData struct {
	dbConfig  *model.DBConfigurationModel
	DBHandler *dbhandler.DBHandler
	Adapter   *adapter.AppHandler
	// AppConfig *model.BaseApplicationConfig
}

func main() {
	// myCoreVariable := &AllModelData{}
	// Load Configuration

	// base, _ := fileconfig.GetBaseAppConfig()
	dbConfig, err := fileconfig.GetDBConfig()
	if err != nil {
		panic(fmt.Sprintf("Error Get DB Config, %v", err))
	}
	// Connect and Migrate DB
	db, err := dbhandler.DBInit(dbConfig)
	if err != nil {
		panic(fmt.Sprintf("Error Init DB, %v", err))
	}

	adapter := &adapter.AppHandler{}

	// open running Server to do Reload Configuration

	adapter.InitializeServer(db)

	myCoreVariable.dbConfig = &dbConfig
	myCoreVariable.DBHandler = db
	myCoreVariable.Adapter = adapter
	// myCoreVariable.AppConfig = &base

	adapter.Run(":4911")

}
