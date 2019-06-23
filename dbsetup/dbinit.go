package dbsetup

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "hasbi"
	password = "dimskii"
	dbname   = "hasbidatabase"
)

type inDB struct {
	database *gorm.DB
}

// DBInit create connection to database
func DBInit() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	baseDb := inDB{database: db}
	if err != nil {
		fmt.Printf("failed to connect Database %v", err)
		panic("failed to connect to database, error")
	}
	fmt.Println("DB Success")
	// migrate table and colomn (setting increment, null, size)
	MigrateResult := db.AutoMigrate(&SystemSettings{}, &Parameters{}, &CaAPIMessages{}, &CaAPITransactions{}, &IsoAPIMessages{})
	if MigrateResult.Error != nil {
		fmt.Printf("failed Migrate Database %v", MigrateResult.Error)
		panic("failed  Migrate Database, error")
	}
	// input index
	DBsetIndex(baseDb.database)
	return db
}

// DBsetIndex for setting Index Database
func DBsetIndex(db *gorm.DB) *gorm.DB {
	db.Model(&CaAPITransactions{}).AddIndex("id_ca_api", "id")
	db.Model(&CaAPITransactions{}).AddIndex("request_id_ca_api", "request_id")
	db.Model(&CaAPITransactions{}).AddIndex("auth_key_ca_api", "auth_key")
	db.Model(&CaAPITransactions{}).AddIndex("customer_number_ca_api", "customer_number")
	db.Model(&CaAPITransactions{}).AddIndex("terminal_id_ca_api", "terminal_id")
	db.Model(&CaAPITransactions{}).AddIndex("rrn_ca_api", "rrn")
	db.Model(&IsoAPIMessages{}).AddIndex("id_iso_api", "id")
	db.Model(&IsoAPIMessages{}).AddIndex("stan_iso_api", "system_trace_audit_number")
	db.Model(&IsoAPIMessages{}).AddIndex("rrn__iso_api", "retrieval_reference_number")
	db.Model(&IsoAPIMessages{}).AddIndex("terminal_id_iso_api", "card_acceptor_terminal_id")
	return nil
}

// DBsetValue is using for set configuration value on DB
func DBsetValue(db *gorm.DB) *gorm.DB {
	result := db.Create(&SystemSettings{Key: "Hasbi", Value: "Meong"})
	if result.Error != nil {
		fmt.Printf("failed input to Database %v", result.Error)
		panic("failed input to database, error")
	}
	return result
}
