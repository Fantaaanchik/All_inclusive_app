package logger

import (
	"all_inclusive_app/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	fileName := config.ConfigureLogger.Path
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50,
		MaxAge:     365,
		MaxBackups: 5,
		Compress:   true,
	})
}
