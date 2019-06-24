package configuration

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// ConfigurationModel file for
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
	
	// put to Model
	dbhost := ConfigurationModel{}
	dbhost.Host = viper.GetString("host")
	port, err := strconv.Atoi(viper.GetString("port"))
	dbhost.Port = port
	dbhost.User = viper.GetString("user")
	dbhost.Password = viper.GetString("password")
	dbhost.Dbname = viper.GetString("dbname")

	// change data now
	viper.Set("host", "127.0.0.1")
	viper.WriteConfig()

	return dbhost
}
