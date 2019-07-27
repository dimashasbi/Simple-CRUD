package fileconfiguration

import (
	"M-GateDBConfig/model"
	"github.com/spf13/viper"
	"strconv"
)

// GetFileConfigInterface interface
type GetFileConfigInterface interface {
	GetDBConfig() (DBConfigurationModel, error)
	GetSimpleConfig() (model.SimpleConfig, error)
	GetIsoMessageConfig() (model.IsoMessageConfig, error)
	GetBaseAppConfig() (model.BaseApplicationConfig, error)
	GetFrontSettingConfig() (model.FrontSettingsConfig, error)
	GetHeaderHTTPMessage() (model.HeaderHTTPMessage, error)
	GetBackSettingConfig() (model.BackSettingsConfig, error)
}

// DBConfigurationModel file for
type DBConfigurationModel struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// GetDBConfig to start get file
func GetDBConfig() (DBConfigurationModel, error) {
	viper.SetConfigType("json")
	viper.AddConfigPath("fileconfiguration/")
	viper.SetConfigName("config")

	dbhost := DBConfigurationModel{}

	err := viper.ReadInConfig()
	if err != nil {
		return dbhost, err
	}

	// put to Model
	dbhost.Host = viper.GetString("db.host")
	port, err := strconv.Atoi(viper.GetString("db.port"))
	dbhost.Port = port
	dbhost.User = viper.GetString("db.user")
	dbhost.Password = viper.GetString("db.password")
	dbhost.Dbname = viper.GetString("db.dbname")
	return dbhost, nil
}

// GetSimpleConfig to get Value from File
func GetSimpleConfig() (model.SimpleConfig, error) {
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("config")

	obj := model.SimpleConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return obj, err
	}

	// put to Model
	obj.Key = viper.GetString("SimpleConfig.key")
	obj.Value = viper.GetString("SimpleConfig.value")

	return obj, nil
}
