package dbservice

// SystemSettings is model for database configuration system
type SystemSettings struct {
	ID          uint `gorm:"primary_key"`
	Registry    string
	IsoRegistry string
}
