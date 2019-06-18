package dbsetup

// SystemSettings is model for database configuration system
type SystemSettings struct {
	// ID    uint `gorm:"primary_key"`
	Key   string
	Value string
}
