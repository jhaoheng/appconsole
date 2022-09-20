package config

import (
	"io"

	"github.com/sirupsen/logrus"
)

type ENV string

const (
	EnvDev  ENV = "dev"
	EnvProd ENV = "prod"
)

var Setting *EnvConfig = &EnvConfig{}

type EnvConfig struct {
	Env ENV `yaml:"env"`
	// log
	LogLevel  logrus.Level     //
	LogFormat logrus.Formatter //
	LogOutput io.Writer        // log 輸出的位置

	// database
	DBSInfo *Database
}

type Database struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func NewConfig(env ENV) *EnvConfig {
	if env == EnvProd {
		Setting = set_prod()
	} else {
		Setting = set_dev()
	}
	return Setting
}
