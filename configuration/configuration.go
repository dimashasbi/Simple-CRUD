package configuration

import (
	// "fmt"
	"strconv"

	"github.com/spf13/viper"
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

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
