package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func set_dev(conf *EnvConfig) {
	conf.LogLevel = logrus.DebugLevel
	conf.LogOutput = os.Stdout
	conf.LogFormat = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	conf.DBSInfo = &Database{
		User:         "",
		Password:     "",
		Host:         "localhost",
		Port:         "3306",
		DatabaseName: "appconsole",
	}
}
