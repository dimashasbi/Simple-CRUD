package postgres

import (
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	parameterRepository struct {
		session *gorm.DB
	}
)

func newParameterRepository(db *gorm.DB) engine.ParameterRepository {
	return &parameterRepository{db}
}

func (p parameterRepository) Insert(m *model.Parameters) error {
	result := p.session.Create(&m)
	if result.Error != nil {
		errStr := "Error input to Parameter Table \n"
		fmt.Printf(errStr, result.Error)
		return errors.Wrapf(result.Error, errStr)
	}
	return nil
}
