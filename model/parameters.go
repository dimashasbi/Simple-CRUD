package model

import "time"

// Parameters for key need to showing in Database
type Parameters struct {
	ID    int       `gorm:"primary_key" json:"-"`
	Key   string    `gorm:"unique;size:60;index:key_param"`
	Value string    `gorm:"size:300"`
	Date  time.Time `json:"-"`
}

// NewParameters create new Parameters
func NewParameters(key, value string) *Parameters {
	return &Parameters{
		Key:   key,
		Value: value,
		Date:  now(),
	}
}
