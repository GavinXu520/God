package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

/*
	logger是一个项目生成日志的工具类

	使用详情可参考logger_test测试类，使用时需要增加logConfig.json的日志配置文件进入项目，具体配置文件请在项目配置文件中声明地址，如下：

	//初始化日志对象（此处在系统配置文件中声明日志文件地址名为logpath）
	logger.SetUpLogger(viper.GetString("logpath"))

	调用方式：
	logger.AlarmLog("back-end_grpc-demo_HealthCheck","system=grpc-demo","系统健康检查")

	此处提供的方法均为warning级别，低于warning级别的记录不在线上环境打印。

	使用示例：
	func TestSetUpLogger(t *testing.T){
	loggerEntry:=&LoggerEntry{
		Level:InfoLevel,
		IsDevelopment:false,
		LogFilePath:"/data/logs/god.log",
		System:"test",
		Namespace:"aura",
	}
	logger:=NewLogger(loggerEntry)
	for i:=0;i<100 ;i++  {
		logger.Info("213254654645")
	}

}

*/

// Level type
type Level uint8

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

type Log struct {
	Timestamp string     `json:"@timestamp"`
	Source    string     `json:"@source"`
	Fields    *LogFields `json:"@fields"`
}
type LogFields struct {
	Interface   string `json:"interface"`
	Method      string `json:"method"`
	Param       string `json:"param"`
	Path        string `json:"path"`
	ProcessTime string `json:"processTime"`
	Result      string `json:"result"`
}
type Entry struct {
	logger      *logrus.Entry
	loggerEntry *LoggerEntry
}

type LoggerEntry struct {
	Level         Level  //输出日志级别,必须
	IsDevelopment bool   //是否是开发环境，开发环境日志输出控制台，否则输出文件,必须
	LogFilePath   string //日志文件绝对路径,必须
	System        string //系统名称,必须
	Namespace     string //命名空间,必须
	HookURL       string //错误级别或者报警日志上报接口，钉钉机器人，可选
	AtPhone       string //设置钉钉@ 某些人的钉钉使用手机号码(多个用","英文逗号分隔)
}

//初始化日志对象
func NewLogger(loggerEntry *LoggerEntry) (entry *Entry) {
	log := &LoggerEntry{
		Level:         loggerEntry.Level,         //日记级别
		IsDevelopment: loggerEntry.IsDevelopment, //是否开发环境
		LogFilePath:   loggerEntry.LogFilePath,   //日志路径
		System:        loggerEntry.System,        //系统名称
		Namespace:     loggerEntry.Namespace,     //命名空间或者项目线名
		HookURL:       loggerEntry.HookURL,
		AtPhone:       loggerEntry.AtPhone,
	}
	logger := log.initLogger()
	entry = &Entry{
		logger:      logger,
		loggerEntry: log,
	}
	return
}

//构建日志对象基本信息
func (this *LoggerEntry) initLogger() (logger *logrus.Entry) {
	if this.IsDevelopment {
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logFile, err := os.OpenFile(this.LogFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			logrus.Fatalf("open file error :%s,原因：%s", this.LogFilePath, err.Error())
			logFile.Close()
		}
		logrus.SetOutput(logFile)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	//级别
	logrus.SetLevel(logrus.Level(this.Level))
	//系统名称
	fullSystemName := this.System + "." + this.Namespace
	logger = logrus.WithFields(logrus.Fields{
		"@source": fullSystemName,
	})
	//hook日志初始化
	if this.HookURL != "" {
		phones := strings.Split(this.AtPhone, ",")
		atPhoneStr := ""
		if phones[0] != "" {
			for _, val := range phones {
				atPhoneStr += "@" + val
			}
		}

		logrus.AddHook(&TalkHook{
			BotURL:         this.HookURL,
			AcceptedLevels: LevelThreshold(logrus.WarnLevel),
			AtPhoneArr:     phones,
			AtPhoneStr:     atPhoneStr,
		})
	}
	return logger
}

func (entry *Entry) AlarmLog(interfaceName string, msg ...interface{}) {
	//格式化msg信息
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02T15:04:05+08:00"),
			"@fields": map[string]interface{}{
				"interface": interfaceName,
				"result":    fieldsMsg,
			},
		}).Warn()
}

func (entry *Entry) Debug(msg ...interface{}) {
	//格式化msg信息
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02T15:04:05+08:00"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Debugln()
}

//
func (entry *Entry) Info(msg ...interface{}) {
	//格式化msg信息
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02T15:04:05+08:00"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Infoln()
}

func (entry *Entry) Warn(msg ...interface{}) {
	//格式化msg信息
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02T15:04:05+08:00"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Warnln()
}

//错误级别的日志报警日志级别输出
func (entry *Entry) Error(msg ...interface{}) {
	//格式化msg信息
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02T15:04:05+08:00"),
			"@fields": map[string]interface{}{
				"interface": entry.loggerEntry.Namespace + "_error",
				"result":    fieldsMsg,
			},
		}).Errorln()
}
