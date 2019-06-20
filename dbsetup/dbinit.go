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
	IndexingResult := db.Model(&CaAPITransactions{}).AddIndex("idx_ID", "id")
	if IndexingResult.Error != nil {
		fmt.Printf("failed input index %v", IndexingResult.Error)
		panic("failed  input idnex, error")
	}
	return db
}

// DBsetValue is using for set configuration value on DB
func DBsetValue(db *gorm.DB) {

	result := db.Create(&SystemSettings{Key: "Hasbi", Value: "Kucing"})
	if result.Error != nil {
		fmt.Printf("failed input to Database %v", result.Error)
		panic("failed input to database, error")
	}
}
