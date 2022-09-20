package config

import (
	"embed"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ENV string

const (
	EnvDev  ENV = "dev"
	EnvProd ENV = "prod"
)

var Setting *EnvConfig = &EnvConfig{}

type EnvConfig struct {
	Env          ENV    `yaml:"env"`
	CommitCode   string `yaml:"commit_code"`
	BuildVersion string `yaml:"version"`
	BuildTime    string
	// log
	LogLevel  logrus.Level     //
	LogFormat logrus.Formatter //
	LogOutput io.Writer        // log 輸出的位置

	// database
	DBSInfo *Database

	//
	Resource *embed.FS
}

type Database struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func NewConfig(yamldata []byte, resource *embed.FS) *EnvConfig {
	if err := yaml.Unmarshal(yamldata, Setting); err != nil {
		panic(err)
	}
	Setting.Resource = resource
	Setting.BuildTime = time.Now().Format("2006-01-02 15:04:05")

	if Setting.Env == EnvProd {
		set_prod(Setting)
	} else {
		set_dev(Setting)
	}
	return Setting
}

func (c *EnvConfig) Show() {
	logrus.Infof("env: %v", Setting.Env)
	logrus.Infof("commit_code: %v", Setting.CommitCode)
	logrus.Infof("version: %v", Setting.BuildVersion)
	logrus.Infof("build_time: %v", Setting.BuildTime)
}
