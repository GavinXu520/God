package log

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestSetUpLogger(t *testing.T) {
	loggerEntry := &LoggerEntry{
		Level:         InfoLevel,
		IsDevelopment: true,
		LogFilePath:   "/data/logs/god.log",
		System:        "test",
		Namespace:     "aura",
		//HookURL:       "https://oapi.dingtalk.com/robot/send?access_token=xxx",
	}
	logger := NewLogger(loggerEntry)
	for i := 0; i < 1; i++ {
		logger.Warn("我要warn级别开始报错了", "awedrwerd")
		logger.Info("我要info级别开始报错了", "awedrwerd")
		logger.Error("我要error开始报错了")
		logger.AlarmLog("postJson请求链接地址:", "http://cardcenter.maizuo.com/getCardInfo.htm", `{"cardNum":"20191203"}`)
	}
}

func TestHttp(t *testing.T) {
	str := `
	{
	    "msgtype": "text",
	    "text": {
		"content": "吃饭了"
	    },
	    "at": {
		"atMobiles": [
		    "18825224373"
		],
		"isAtAll": false
	    }
	}
	`
	talk_url := "https://oapi.dingtalk.com/robot/send?access_token=c05d504a7f02940e89333f1cab32d513479f0a659f2db685752c52779d956e37"
	body := strings.NewReader(str)
	resp, _ := http.Post(talk_url, "application/json", body)
	defer resp.Body.Close()
}

func TestNewLogger(t *testing.T) {
	headerByte, _ := base64.StdEncoding.DecodeString("")
	header := string(headerByte)
	fmt.Println(header)
}
