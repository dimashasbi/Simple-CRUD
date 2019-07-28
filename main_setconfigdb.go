package main

import (
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/fileconfiguration"
	"M-GateDBConfig/handler"
	"M-GateDBConfig/model"
	"M-GateDBConfig/mylog"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllModelData that have Global Value
type AllModelData struct {
	dbConfig     *model.DBConfigurationModel
	DBHandler    *dbhandler.DBHandler
	AppHandler   *handler.AppHandler
	ModAppConfig *model.BaseApplicationConfig
}

func main() {
	myCoreVariable := &AllModelData{}
	// Load Configuration
	base, err1 := fileconfiguration.GetBaseAppConfig()
	dbConfig, err2 := fileconfiguration.GetDBConfig()
	err3 := mylog.InitLog(base)
	checkErr(err1, dbConfig)
	checkErr(err2, dbConfig)
	checkErr(err3, "Init Log")
	// Connect and Migrate DB
	db, err4 := dbhandler.DBInit(dbConfig)
	checkErr(err4, db)

	handler := &handler.AppHandler{}
	// open running Server to do Reload Configuration
	handler.InitializeServer(db)

	myCoreVariable.dbConfig = &dbConfig
	myCoreVariable.DBHandler = db
	myCoreVariable.AppHandler = handler
	myCoreVariable.ModAppConfig = &base

	handler.Run(":4911")

}

func checkErr(err error, data interface{}) {
	if err != nil {
		mylog.FatalError(err)
		// panic(err)

	}
	mylog.Debug("Main Application ", data)
}
