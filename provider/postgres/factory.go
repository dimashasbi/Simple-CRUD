package postgres

import (
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	storageFactory struct {
		db *gorm.DB
	}
)

// NewStorage creates a new instance of this mongodb storage factory
func NewStorage(data model.DBConfigurationModel) engine.StorageFactory {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.Dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect Database")
	}
	fmt.Println("Database Connect")
	// migrate table and colomn (setting increment, null, size)
	MigrateResult := db.AutoMigrate(&model.Registry{}, &model.Parameters{})
	if MigrateResult.Error != nil {
		panic("failed  Migrate Database")
	}
	fmt.Println("Database Migrate Success")

	// input index
	dbsetIndex(db)

	// done
	fmt.Println("Database Init Success")
	return &storageFactory{db}

}

func (c *storageFactory) NewParameterRepository() engine.ParameterRepository {
	return newParameterRepository(c.db)
}

// dbsetIndex to Set Index
func dbsetIndex(db *gorm.DB) {
	// db.Model(&model.Parameters{}).AddIndex("key_param", "key")
	// db.Model(&model.Registry{}).AddIndex("key_reg", "key")
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
