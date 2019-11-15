package main

import (
	"M-GateDBConfig/adapter"
	"M-GateDBConfig/app/fileconfig"
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"
	"M-GateDBConfig/provider/postgres"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type (
	TestHandler struct {
		dbHandler *gorm.DB
	}
)

var (
	testDBValue = model.DBConfigurationModel{
		Dbname:   "hasbidatabase",
		Password: "dimskii",
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "hasbi",
	}
)

func TestGetDBConfig(T *testing.T) {
	// Test if value JSON is not same

	// // create file config
	// EditDBconfValue(testValue)

	// test - get fileconfig
	db := fileconfig.GetDBConfig()

	checkdbconfvalue(testDBValue, db, T)

	// return real value for Development
	returnRealDBconfValue()

}

func InitializeApps() {
	// Load Configuration
	dbConfig := fileconfig.GetDBConfig()
	// Connect and Migrate DB
	db := postgres.NewStorage(dbConfig)
	// Prepare Engine for Use Case Logic
	eng := engine.NewEngine(db)
	// Start application
	adapter := adapter.Handler{}
	adapter.InitializeServer(eng)
	// adapter.Run(":4911")
}

func TestNewStorage(T *testing.T) {
	// 1. Test Table Existing or Not after Migration

	// DB Migration first
	postgres.NewStorage(testDBValue)

	// Test Table Existing or Not
	// connect DB
	data := testDBValue
	db := InitDB(data)
	// checkTable exist or not
	result := db.Exec("SELECT * FROM REGRISTRIES")
	if result.Error != nil {
		T.Errorf("No Table Exists %+v\n", result.Error)
	}
}

func TestInsertParameterToDB(T *testing.T) {
	// Initialize Application First
	InitializeApps()

	// Test Insert Parameter to Database

	// Test unique KEY

	// Test no Table

	// Test no Connection of Database

}

func returnRealDBconfValue() {
	fileconfig.CreateDBConfig()
}

func EditDBconfValue(nan model.DBConfigurationModel) {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath("./")
	v.SetDefault("db", nan)
	v.WriteConfigAs("appconfig.json")
}

func checkdbconfvalue(expected, actual model.DBConfigurationModel, t *testing.T) {
	if expected != actual {
		t.Errorf("Not Same Configuration %+v. Got %+v\n", expected, actual)
	}
}


func InitDB(data model.DBConfigurationModel) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.Dbname)
	dba, _ := gorm.Open("postgres", psqlInfo)
	return dba
}