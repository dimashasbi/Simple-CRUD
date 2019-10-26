package fileconfig

import (
	"M-GateDBConfig/model"
	"strconv"

	"github.com/spf13/viper"
)

// GetFileConfigInterface interface
type GetFileConfigInterface interface {
	GetDBConfig() (model.DBConfigurationModel, error)
	GetSimpleConfig() (model.SimpleConfig, error)
	GetIsoMessageConfig() (model.IsoMessageConfig, error)
	// GetBaseAppConfig() (model.BaseApplicationConfig, model.DBConfigurationModel, error)
	GetFrontSettingConfig() (model.FrontSettingsConfig, error)
	GetHeaderHTTPMessage() (model.HeaderHTTPMessage, error)
	GetBackSettingConfig() (model.BackSettingsConfig, error)
}

// GetDBConfig to start get file
func GetDBConfig() (model.DBConfigurationModel, error) {
	// mylog.MnDebug("Load DB Configuration ")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("appconfig")

	dbhost := model.DBConfigurationModel{}

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

// // GetBaseAppConfig use for get Base Application Configuration
// func GetBaseAppConfig() (model.BaseApplicationConfig, error) {
// 	// mylog.MnDebug("Load BaseApp Config")
// 	viper.SetConfigType("json")
// 	viper.AddConfigPath("./")
// 	viper.SetConfigName("appconfig")

// 	obj := model.BaseApplicationConfig{}

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return obj, err
// 	}

// 	// put to Model
// 	obj.CaName = viper.GetString("BaseApplicationConfig.CaName")
// 	obj.LogFolder = viper.GetString("BaseApplicationConfig.LogFolder")
// 	obj.ActivateSVA = viper.GetString("BaseApplicationConfig.ActivateSVA")
// 	obj.PlnDirect = viper.GetString("BaseApplicationConfig.PlnDirect")

// 	return obj, nil
// }

// // GetSimpleConfig to get Value from File
// func GetSimpleConfig() (model.SimpleConfig, error) {
// 	// mylog.MnDebug("Load Simple Config")
// 	viper.SetConfigType("json")
// 	viper.AddConfigPath("configuration/")
// 	viper.SetConfigName("setdbvalue")

// 	obj := model.SimpleConfig{}

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return obj, err
// 	}

// 	// put to Model
// 	obj.Key = viper.GetString("SimpleConfig.key")
// 	obj.Value = viper.GetString("SimpleConfig.value")

// 	return obj, nil
// }