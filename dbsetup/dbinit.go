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
	db.AutoMigrate(SystemSettings{})
	return db
}

// set configuration
func DBsetValue(db *gorm.DB) {

	result := db.Create(&SystemSettings{Key: "Hasbi", Value: "Kucing"})
	if result.Error != nil {
		fmt.Printf("failed input to Database %v", result.Error)
		panic("failed input to database, error")
	}
}
