package configuration

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tkanos/gonfig"
)

// Configuration Model file
type Configuration struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func getConfig(ya bool) bool {
	//filename is the path to the json config file
	err := gonfig.GetConf("config/config.json", &configuration)
	fmt.Printf("configuration port", Configuration)
	return true
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
