package fileconfig

import (
	"M-GateDBConfig/model"

	"github.com/spf13/viper"
)

// CreateDBConfig to create File config
func CreateDBConfig() {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath("./")
	dbhost := model.DBConfigurationModel{
		Dbname:   "hasbidatabase",
		Password: "dimskii",
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "hasbi",
	}

	// err := v.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	v.SetDefault("db", dbhost)
	v.WriteConfigAs("appconfig.json")
}

// GetDBConfig to start get file
func GetDBConfig() model.DBConfigurationModel {
	// mylog.MnDebug("Load DB Configuration ")
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath("./")
	v.SetConfigName("appconfig")

	dbhost := model.DBConfigurationModel{}

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// put to Model
	dbhost.Host = v.GetString("db.host")
	// port, err := strconv.Atoi(viper.GetString("db.port"))
	dbhost.Port = v.GetInt("db.port")
	dbhost.User = v.GetString("db.user")
	dbhost.Password = v.GetString("db.password")
	dbhost.Dbname = v.GetString("db.dbname")
	return dbhost
}
