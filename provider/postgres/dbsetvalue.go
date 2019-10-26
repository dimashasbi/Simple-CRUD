package postgres

import (
	"M-GateDBConfig/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// IDBGetSet interface
type IDBGetSet interface {
	SetParameter(data model.Parameters) error
	// SetSystemSettings(data model.SystemSettings) error

	GetParameter() ([]model.Parameters, error)
	// GetSystemSettings() (model.SystemSettings, error)
}

// DBHandler for db structure
type DBHandler struct {
	database *gorm.DB
}

// SetParameter is using for set configuration value on DB
func (db *DBHandler) SetParameter(data model.Parameters) error {
	// put model to Database
	result := db.database.Create(&data)
	if result.Error != nil {
		errStr := "Error input to Parameter Table \n"
		fmt.Printf(errStr, result.Error)
		return errors.Wrapf(result.Error, errStr)
	}
	return nil
}

// GetParameter is using for set configuration value on DB
func (db *DBHandler) GetParameter() ([]model.Parameters, error) {
	// put model to Database
	var data []model.Parameters
	result := db.database.Find(&data)
	// result := db.database.First(&data)
	if result.Error != nil {
		errStr := "Error Get Value from Parameter Table \n"
		return data, errors.Wrapf(result.Error, errStr)
	}
	return data, nil
}

// // SetSystemSettings is using for set configuration value on DB
// func (db *DBHandler) SetSystemSettings(data model.SystemSettings) error {
// 	// put model to Database
// 	result := db.database.Create(&data)
// 	if result.Error != nil {
// 		errStr := "Error input to SystemSettings Table \n"
// 		fmt.Printf(errStr, result.Error)
// 		// panic(result.Error)
// 		return result.Error
// 	}
// 	return nil
// }
