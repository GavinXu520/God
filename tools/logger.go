package tools

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

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
	Level         Level  // Logger level, required
	IsDevelopment bool   // Is development?，the log print out on console if it be development, Otherwise, print out in file. required
	LogFilePath   string // The log file absolute path, required
	System        string // This service name, required
}

func NewLogger(loggerEntry *LoggerEntry) (entry *Entry) {
	log := &LoggerEntry{
		Level:         loggerEntry.Level,         // the logger level
		IsDevelopment: loggerEntry.IsDevelopment, // the env is development ?
		LogFilePath:   loggerEntry.LogFilePath,   // the log file path
		System:        loggerEntry.System,        // service sys name
	}
	logger := log.initLogger()
	entry = &Entry{
		logger:      logger,
		loggerEntry: log,
	}
	return
}

// Build the logger basic instance
func (this *LoggerEntry) initLogger() (logger *logrus.Entry) {
	if this.IsDevelopment {
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logFile, err := os.OpenFile(this.LogFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			logrus.Fatalf("open file error :%s, the reason is：%s", this.LogFilePath, err.Error())
			logFile.Close()
		}
		logrus.SetOutput(logFile)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetLevel(logrus.Level(this.Level))

	fullSystemName := this.System
	logger = logrus.WithFields(logrus.Fields{
		"@sysName": fullSystemName,
	})
	return logger
}

func (entry *Entry) Debug(msg ...interface{}) {
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Debugln()
}

func (entry *Entry) Info(msg ...interface{}) {
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Infoln()
}

func (entry *Entry) Warn(msg ...interface{}) {
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"result": fieldsMsg,
			},
		}).Warnln()
}

func (entry *Entry) Error(msg ...interface{}) {
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"interface": entry.loggerEntry.System + "_error",
				"result":    fieldsMsg,
			},
		}).Errorln()
}

func (entry *Entry) AlarmLog(interfaceName string, msg ...interface{}) {
	// format the msg ...
	msgSprint := fmt.Sprintln(msg...)
	fieldsMsg := msgSprint[:len(msgSprint)-1]
	entry.logger.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"interface": interfaceName,
				"result":    fieldsMsg,
			},
		}).Warn()
}
