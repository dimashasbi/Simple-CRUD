package dbhandler

import (
	"fmt"
)

// // DBHandler use for interface every object use this package
// type DBHandler interface {
// 	SetDBValue(db *gorm.DB, val string) error
// 	GetDBValue(db *gorm.DB, num int) (string, error)
// }

// SetDBValue interface
// type SetDBValue interface {
// 	SetParameter(key string, val string) error
// 	SetSystemSettings(key string, val string) error
// 	SetCaAPIMessages(mod *CaAPIMessages) error
// 	SetCaAPITransactions(mod *CaAPITransactions) error
// 	SetIsoAPIMessages(mod *IsoAPIMessages) error
// }

// SetParameter is using for set configuration value on DB
func (db *DBHandler) SetParameter(data *Parameters) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to Parameter Table \n"
		fmt.Printf(errStr, result.Error)
		// panic(result.Error)
		return result.Error
	}
	return nil
}

// SetSystemSettings is using for set configuration value on DB
func (db *DBHandler) SetSystemSettings(key string, val string) error {
	// put model to Database
	mod1 := SystemSettings{Key: key, Value: val}
	result := db.database.Create(mod1)
	if result.Error != nil {
		errStr := "Error input to System Settings Table"
		fmt.Printf(errStr)
		return result.Error
	}
	return nil
}

// SetCaAPIMessages is using for set configuration value on DB
func (db *DBHandler) SetCaAPIMessages(key string, val string) error {
	// put model to Database
	mod1 := SystemSettings{Key: key, Value: val}
	result := db.database.Create(mod1)
	if result.Error != nil {
		errStr := "Error input to CaAPIMessages Table"
		fmt.Printf(errStr)
		return result.Error
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error in Database")
	}
}
