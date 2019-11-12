package model

import (
	"time"
)

// Users is model for database configuration system
type Users struct {
	ID                 int32  `gorm:"size:30;index:users_id;unique"`
	UserName           string `gorm:"size:30;index:users_name;unique"`
	Password           string `gorm:"size:50;index:users_pass"`
	UserFullName       string `gorm:"size:50"`
	Email              string `gorm:"size:50"`
	RoleID             int    `gorm:"size:1;index:users_roleid"`
	LastLoginTime      time.Time
	LoginFailCount     int `gorm:"size:1;index:users_lgnfailcnt"`
	LastChangePassTime time.Time
	MustChangePass     bool `gorm:"index:users_muschgpas"`
	DisableReasonCode  int  `gorm:"size:1;index:users_disbresncd"`
	CreatedTime        time.Time
	Disabled           bool `gorm:"index:users_disbl"`
	Active             bool `gorm:"index:users_actv"`
}

// NewUsers create new Registry
func NewUsers(key, value string) *Registry {
	return &Registry{
		Key:   key,
		Value: value,
		Date:  now(),
	}
}
