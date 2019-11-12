package postgres

import (
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	systemSettingsRepository struct {
		session *gorm.DB
	}
)

func newSystemSettingsRepository(db *gorm.DB) engine.SystemSettingsRepository {
	return &systemSettingsRepository{db}
}

func (s systemSettingsRepository) Select(k *model.Registry) (*model.Registry, error) {
	result := s.session.Where("KEY = ?", k.Key).First(&k)
	if result.Error != nil {
		return k, errors.Errorf("Error Select a Registry : %v", result.Error)
	}
	return k, nil
}

func (s systemSettingsRepository) Insert(k *model.Registry) error {
	result := s.session.Create(&k)
	if result.Error != nil {
		return errors.Errorf("Error Insert Registries : %v", result.Error)
	}
	return nil
}

func (s systemSettingsRepository) Update(k *model.Registry) error {
	result := s.session.Model(&k).Update("value", k.Value)
	if result.Error != nil {
		return errors.Errorf("Error Update Registries : %v", result.Error)
	}
	return nil
}
