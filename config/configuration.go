package config

import "os"

type databaseConfig struct {
	Username     string
	Password     string
	Address      string
	DatabaseName string
}

func LoadDatabaseConfiguration() (databaseConfig, error) {
	var result databaseConfig
	result.Username = "root"
	result.Password = "password"
	result.Address = "127.0.0.1"
	result.DatabaseName = "orderValidator"
	return result, nil

}

func LoadPort() string {
	return os.Getenv("PORT")
	// return "8080"
}
