package model

import (
	"time"
)

// Registry is model for database configuration system
type Registry struct {
	ID    uint   `gorm:"primary_key"`
	Key   string `gorm:"size:40"`
	Value string `gorm:"size:2000"`
	Date  time.Time
}

// NewRegistry create new Registry
func NewRegistry(key, value string) *Registry {
	return &Registry{
		Key:   key,
		Value: value,
		Date:  now(),
	}
}
