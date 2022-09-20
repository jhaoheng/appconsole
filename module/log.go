package module

import (
	"appconsole/config"

	"github.com/sirupsen/logrus"
)

func SetLog(conf *config.EnvConfig) {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(conf.LogFormat)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(conf.LogOutput)

	// Only log the warning severity or above.
	logrus.SetLevel(conf.LogLevel)

	//
	logrus.SetReportCaller(true)
}
