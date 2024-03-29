package main

import (
	"M-GateDBConfig/adapter"
	"M-GateDBConfig/provider/fileconfig"
	"M-GateDBConfig/engine"

	"M-GateDBConfig/provider/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// // create file config
	// fileconfig.CreateDBConfig()

	// Load Configuration
	dbConfig := fileconfig.GetDBConfig()

	// Connect and Migrate DB
	db := postgres.NewStorage(dbConfig)
	// Prepare Engine for Use Case Logic
	eng := engine.NewEngine(db)
	// Start application
	adapter := adapter.Handler{}
	adapter.InitializeServer(eng)
	adapter.Run(":4911")

}
