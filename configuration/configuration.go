package configuration

import (
	"M-GateDBConfig/model"
	"github.com/spf13/viper"
	"strconv"
)

// DBConfigurationModel file for
type DBConfigurationModel struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// GetDBConfig to start get file
func GetDBConfig() DBConfigurationModel {
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	checkErr(err)

	// put to Model
	dbhost := DBConfigurationModel{}
	dbhost.Host = viper.GetString("db.host")
	port, err := strconv.Atoi(viper.GetString("db.port"))
	dbhost.Port = port
	dbhost.User = viper.GetString("db.user")
	dbhost.Password = viper.GetString("db.password")
	dbhost.Dbname = viper.GetString("db.dbname")

	return dbhost
}

// GetParamValue to get Value from File
func GetParamValue() *model.SimpleConfig {
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	checkErr(err)

	// put to Model
	obj := &model.SimpleConfig{}
	obj.Key = viper.GetString("param.key")
	obj.Value = viper.GetString("param.value")

	return obj
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
