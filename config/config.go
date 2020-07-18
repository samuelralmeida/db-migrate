package config

import (
	"github.com/spf13/viper"
)

type dbConfig struct {
	Host     string
	Port     string
	Username string
	Database string
	Password string
	SSLMode  string
}

// DBConfig is a struct with data to database connect
var DBConfig dbConfig

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	DBConfig = dbConfig{}

	DBConfig.Database = viper.GetString("database")
	DBConfig.Host = viper.GetString("host")
	DBConfig.Password = viper.GetString("password")
	DBConfig.Port = viper.GetString("port")
	DBConfig.Username = viper.GetString("username")
	DBConfig.SSLMode = viper.GetString("sslmode")
}
