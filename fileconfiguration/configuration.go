package fileconfiguration

import (
	"M-GateDBConfig/model"
	"M-GateDBConfig/mylog"
	"github.com/spf13/viper"
	"strconv"
)

// GetFileConfigInterface interface
type GetFileConfigInterface interface {
	GetDBConfig() (model.DBConfigurationModel, error)
	GetSimpleConfig() (model.SimpleConfig, error)
	GetIsoMessageConfig() (model.IsoMessageConfig, error)
	GetBaseAppConfig() (model.BaseApplicationConfig, error)
	GetFrontSettingConfig() (model.FrontSettingsConfig, error)
	GetBackSettingConfig() (model.BackSettingsConfig, error)
}

// GetDBConfig to start get file
func GetDBConfig() (model.DBConfigurationModel, error) {
	mylog.MnDebug("Load DB Configuration ")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("appconfig")

	dbhost := model.DBConfigurationModel{}

	err := viper.ReadInConfig()
	if err != nil {
		return dbhost, err
		// format := "Error Read File Configuration \n"
		// sol := "Check Your Configuration File Path"
		// return dbhost, errors.Wrapf(err, format, sol)
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

// GetBaseAppConfig use for get Base Application Configuration
func GetBaseAppConfig() (model.BaseApplicationConfig, error) {
	mylog.MnDebug("Load BaseApp Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("appconfig")

	obj := model.BaseApplicationConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return obj, err
	}

	// put to Model
	obj.CaName = viper.GetString("BaseApplicationConfig.CaName")
	obj.LogFolder = viper.GetString("BaseApplicationConfig.LogFolder")
	obj.ActivateSVA = viper.GetString("BaseApplicationConfig.ActivateSVA")
	obj.PlnDirect = viper.GetString("BaseApplicationConfig.PlnDirect")
	obj.RunningPort = viper.GetString("BaseApplicationConfig.RunningPort")

	return obj, nil
}

// GetSimpleConfig to get Value from File
func GetSimpleConfig() (model.SimpleConfig, error) {
	mylog.MnDebug("Load Simple Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("setdbvalue")

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

// GetIsoMessageConfig to get Value from File
func GetIsoMessageConfig() (model.IsoMessageConfig, error) {
	mylog.MnDebug("Load Simple Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("setdbvalue")

	obj := model.IsoMessageConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return obj, err
	}

	// put to Model
	obj.BitActive = viper.GetString("IsoMessageConfig.BitActive")
	obj.MerchantType = viper.GetString("IsoMessageConfig.MerchantType")
	obj.PosEntryMode = viper.GetString("IsoMessageConfig.PosEntryMode")
	obj.AcquiringInstitutionID = viper.GetString("IsoMessageConfig.AcquiringInstitutionID")
	obj.TerminalID = viper.GetString("IsoMessageConfig.TerminalID")
	obj.CardAcceptorID = viper.GetString("IsoMessageConfig.CardAcceptorID")
	obj.CardAcceptorName = viper.GetString("IsoMessageConfig.CardAcceptorName")
	obj.ForwardingInstitutionID = viper.GetString("IsoMessageConfig.ForwardingInstitutionID")

	return obj, nil
}

// GetFrontSettingConfig to get Value from File
func GetFrontSettingConfig() (model.FrontSettingsConfig, error) {
	mylog.MnDebug("Load Simple Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("setdbvalue")

	obj := model.FrontSettingsConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return obj, err
	}

	// put to Model
	obj.HTTPListeningPort = viper.GetString("FrontSettingsConfig.HTTPListeningPort")
	obj.HeaderMessage.ClientID = viper.GetString("FrontSettingsConfig.ClientID")
	obj.HeaderMessage.PassKey = viper.GetString("FrontSettingsConfig.PassKey")
	obj.HeaderMessage.SecretKey = viper.GetString("FrontSettingsConfig.SecretKey")

	return obj, nil
}

// GetBackSettingConfig to get Value from File
func GetBackSettingConfig() (model.BackSettingsConfig, error) {
	mylog.MnDebug("Load Simple Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("setdbvalue")

	obj := model.BackSettingsConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return obj, err
	}

	// put to Model
	obj.URLJavaMPAY = viper.GetString("BackSettingsConfig.URLJavaMPAY")
	obj.URLJavaSVA = viper.GetString("BackSettingsConfig.URLJavaSVA")

	return obj, nil
}
