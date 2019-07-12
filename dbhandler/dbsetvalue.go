package dbsetup

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBHandler interface {
	SetDBValue
	GetDBValue
}

type SetDBValue struct {
	DBsetBackSettingsConfig()
} 


// DBsetBackSettingsConfig is using for set configuration value on DB
func DBsetBackSettingsConfig(db *gorm.DB, val string) (bool, error) {

	// put model to Database
	mod1 := SystemSettings{Key: "SETTINGS1", Value: val}
	result := db.Create(mod1)
	if result.Error != nil {
		fmt.Printf("Error input Back Settings Value")
		return false, nil
	}
	return true, nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error in Database")
	}
}
