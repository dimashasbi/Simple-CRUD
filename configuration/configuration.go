package configuration

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// Configuration Model file
type ConfigurationModel struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// GetConfig to start get file
func GetConfig() ConfigurationModel {
	viper.SetConfigType("json")
	viper.AddConfigPath("configuration/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("ga bisa bacaaa : %v ", err)
	}
	fmt.Printf("bisa dibaca  %v \n ", viper.GetString("port"))

	// put to Model
	dbhost := ConfigurationModel{}
	dbhost.Host = viper.GetString("host")
	port, err := strconv.Atoi(viper.GetString("port"))
	dbhost.Port = port
	dbhost.User = viper.GetString("user")
	dbhost.Password = viper.GetString("password")
	dbhost.Dbname = viper.GetString("dbname")

	return dbhost
}
