package dbservice

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
