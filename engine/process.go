package engine

import (
	"net/http"
	"M-GateDBConfig/model"
	"M-GateDBConfig/provider/dbhandler"
	"fmt"
)

// Idatabase
type IDatabase struct {
	IDB   dbhandler.IDBGetSet
}

// Parameter for output show Table Parameter
type Parameter struct {
	Param []model.Parameters `json:"Parameter"`
}

// GetParam to input value to DB
func GetParam(IDB dbhandler.IDBGetSet) (Parameter, error) {
	output := Parameter{}
	data, err := IDB.GetParameter()
	if err != nil {
		fmt.Printf("Error get parameter %v \n", err)
		return output, err
	}
	output.Param = data

	return output, nil
}

// SetSimpleConfig to input value to DB
func (db IDatabase)SetSimpleConfig(m model.SimpleConfig) (Parameter, error) {
	mod := model.SimpleConfig{}

	

	data, err := IDB.GetParameter()
	if err != nil {
		fmt.Printf("Error get parameter %v \n", err)
		return output, err
	}
	output.Param = data

	return output, nil
}
