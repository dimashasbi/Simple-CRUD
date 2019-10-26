package postgres

import (
	"M-GateDBConfig/model"

	"fmt"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit(data model.DBConfigurationModel) (*DBHandler, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.Dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	baseDb := &DBHandler{database: db}
	if err != nil {

		// fmt.Printf("failed to connect Database %v", err)
		// panic("failed to connect Database")
		return nil, err
	}
	fmt.Println("Database Connect")

	// migrate table and colomn (setting increment, null, size)
	MigrateResult := db.AutoMigrate(&model.Registry{}, &model.Parameters{})
	if MigrateResult.Error != nil {

		fmt.Printf("failed Migrate Database %v", MigrateResult.Error)
		// panic("failed  Migrate Database")
		return nil, MigrateResult.Error
	}
	fmt.Println("Database Migrate Success")

	// input index
	DBsetIndex(baseDb.database)
	fmt.Println("Database Init Success")
	return baseDb, nil
}

// DBsetIndex for setting Index Database
func DBsetIndex(db *gorm.DB) {
	// db.Model(&model.CaAPITransactions{}).AddIndex("id_ca_api", "id")
	// db.Model(&model.CaAPITransactions{}).AddIndex("request_id_ca_api", "request_id")
	// db.Model(&model.CaAPITransactions{}).AddIndex("auth_key_ca_api", "auth_key")
	// db.Model(&model.CaAPITransactions{}).AddIndex("customer_number_ca_api", "customer_number")
	// db.Model(&model.CaAPITransactions{}).AddIndex("terminal_id_ca_api", "terminal_id")
	// db.Model(&model.CaAPITransactions{}).AddIndex("rrn_ca_api", "rrn")
	// db.Model(&model.IsoAPIMessages{}).AddIndex("id_iso_api", "id")
	// db.Model(&model.IsoAPIMessages{}).AddIndex("stan_iso_api", "system_trace_audit_number")
	// db.Model(&model.IsoAPIMessages{}).AddIndex("rrn__iso_api", "retrieval_reference_number")
	// db.Model(&model.IsoAPIMessages{}).AddIndex("terminal_id_iso_api", "card_acceptor_terminal_id")
}
