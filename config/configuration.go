package config

import "os"

func LoadDatabaseConfiguration() string {
	return os.Getenv("DATABASE_URL")
}

func LoadPort() string {
	return os.Getenv("PORT")
}

// HOST = "ec2-54-172-169-87.compute-1.amazonaws.com"
// DatabaseName = "dcmsaj87g5dsfj"
// USER = "rwgnzabvwngdls"
// Port = "5432"
// Password = "f910bbc7451e346bff935f7595f6f1d06394944460eff6f183801248d95bb91e"
// URI = "    postgres://rwgnzabvwngdls:f910bbc7451e346bff935f7595f6f1d06394944460eff6f183801248d95bb91e@ec2-54-172-169-87.compute-1.amazonaws.com:5432/dcmsaj87g5dsfj"
