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
		logger.Error("我要error开始报错了")
		logger.AlarmLog("postJson请求链接地址:", "www.god.com", `{"cardNum":"20191203"}`)
	}
}
