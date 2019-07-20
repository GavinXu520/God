package common

import (
	"God/tools"

	"github.com/spf13/viper"
)

var (
	Logger *tools.Entry
)

func SetUpLogger() {
	isDevelopment := viper.GetBool("isDevelopment")
	logFilePath := viper.GetString("log.path")
	system := viper.GetString("name")

	logLevel := viper.GetString("log.level")

	loggerEntry := &tools.LoggerEntry{
		Level:         parseLevel(logLevel),
		IsDevelopment: isDevelopment,
		LogFilePath:   logFilePath,
		System:        system,
	}
	Logger = tools.NewLogger(loggerEntry)
}

func AlarmLog(interfaceName string, param interface{}, other interface{}) {
	Logger.AlarmLog(interfaceName, param, other)
}

func parseLevel(level string) tools.Level {

	var l tools.Level

	switch level {
	case "panic":
		l = tools.PanicLevel
	case "fatal":
		l = tools.FatalLevel
	case "error":
		l = tools.ErrorLevel
	case "warn":
		l = tools.WarnLevel
	case "info":
		l = tools.InfoLevel
	case "debug":
		l = tools.DebugLevel
	default:
		l = tools.DebugLevel
	}
	return l
}
