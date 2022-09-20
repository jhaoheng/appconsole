package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func set_prod() *EnvConfig {
	return &EnvConfig{
		Env:       "prod",
		LogLevel:  logrus.InfoLevel,
		LogOutput: SetLogPosition("./log"),
		LogFormat: &logrus.JSONFormatter{
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

/*
- 若儲存在檔案中, 設定旋轉 log
- 常駐程式, 超過檔案大小會自動旋轉
- 手動判斷旋轉請參考文件
*/
func SetLogPosition(filename string) *lumberjack.Logger {
	l := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     30,    //days
		Compress:   false, // disabled by default
	}
	return l
}
