package dbhandler

import (
	"M-GateDBConfig/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// // DBHandler use for interface every object use this package
// type DBHandler interface {
// 	SetDBValue(db *gorm.DB, val string) error
// 	GetDBValue(db *gorm.DB, num int) (string, error)
// }

// ISetDBValue interface
type ISetDBValue interface {
	SetParameter(data *model.Parameters) error
	SetSystemSettings(data *model.SystemSettings) error
	SetCaAPIMessages(mod *model.CaAPIMessages) error
	SetCaAPITransactions(mod *model.CaAPITransactions) error
	SetIsoAPIMessages(mod *model.IsoAPIMessages) error
}

// DBHandler for db structure
type DBHandler struct {
	database *gorm.DB
}

// SetParameter is using for set configuration value on DB
func (db *DBHandler) SetParameter(data *model.Parameters) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to Parameter Table \n"
		fmt.Printf(errStr, result.Error)
		return errors.Wrapf(result.Error, errStr, "Apaa yaa")
	}
	return nil
}

// SetSystemSettings is using for set configuration value on DB
func (db *DBHandler) SetSystemSettings(data *model.SystemSettings) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to SystemSettings Table \n"
		fmt.Printf(errStr, result.Error)
		// panic(result.Error)
		return result.Error
	}
	return nil
}

// SetCaAPIMessages is using for set configuration value on DB
func (db *DBHandler) SetCaAPIMessages(data *model.CaAPIMessages) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to CaAPIMessages Table \n"
		fmt.Printf(errStr, result.Error)
		// panic(result.Error)
		return result.Error
	}
	return nil
}

// SetCaAPITransactions is using for set configuration value on DB
func (db *DBHandler) SetCaAPITransactions(data *model.CaAPITransactions) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to CaAPITransactions Table \n"
		fmt.Printf(errStr, result.Error)
		// panic(result.Error)
		return result.Error
	}
	return nil
}

// SetIsoAPIMessages is using for set configuration value on DB
func (db *DBHandler) SetIsoAPIMessages(data *model.IsoAPIMessages) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to IsoAPIMessages Table \n"
		fmt.Printf(errStr, result.Error)
		// panic(result.Error)
		return result.Error
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error in Database")
	}
}
