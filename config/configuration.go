package config

import "os"

type databaseConfig struct {
	Username     string
	Password     string
	Address      string
	DatabaseName string
}

// func LoadDatabaseConfiguration() (databaseConfig, error) {
// 	var result databaseConfig
// 	result.Username = "root"
// 	result.Password = "password"
// 	result.Address = "127.0.0.1"
// 	result.DatabaseName = "orderValidator"
// 	return result, nil

// }

func LoadPort() string {
	return os.Getenv("PORT")
	// return "8080"
}

// HOST = "ec2-54-172-169-87.compute-1.amazonaws.com"
// DatabaseName = "dcmsaj87g5dsfj"
// USER = "rwgnzabvwngdls"
// Port = "5432"
// Password = "f910bbc7451e346bff935f7595f6f1d06394944460eff6f183801248d95bb91e"
// URI = "    postgres://rwgnzabvwngdls:f910bbc7451e346bff935f7595f6f1d06394944460eff6f183801248d95bb91e@ec2-54-172-169-87.compute-1.amazonaws.com:5432/dcmsaj87g5dsfj"
