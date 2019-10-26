package model

import "time"

// Parameters for key need to showing in Database
type Parameters struct {
	ID    int    `gorm:"primary_key"`
	Key   string `gorm:"size:60"`
	Value string `gorm:"size:300"`
	Date  time.Time
}

// NewParameters create new Parameters
func NewParameters(key, value string) *Parameters {
	return &Parameters{
		Key:   key,
		Value: value,
		Date:  now(),
	}
}
