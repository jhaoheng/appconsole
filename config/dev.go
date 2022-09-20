package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func set_dev() *EnvConfig {
	return &EnvConfig{
		Env:       "dev",
		LogLevel:  logrus.DebugLevel,
		LogOutput: os.Stdout,
		LogFormat: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		//
		DBSInfo: &Database{
			User:         "",
			Password:     "",
			Host:         "localhost",
			Port:         "3306",
			DatabaseName: "appconsole",
		},
	}
}
