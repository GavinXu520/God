package tools

import (
	"testing"
)

func TestSetUpLogger(t *testing.T) {
	loggerEntry := &LoggerEntry{
		Level:         InfoLevel,
		IsDevelopment: true,
		LogFilePath:   "/data/logs/god.log",
		System:        "test",
	}
	logger := NewLogger(loggerEntry)
	for i := 0; i < 1; i++ {
		logger.Warn("I am warn log", "i am ...")
		logger.Info("I am info log", "i am ...")
		logger.Error("I am error log", "i am ...")
		logger.AlarmLog("This is func interface name", "www.god.com", `{"cardNum":"20191203"}`)
	}
}
