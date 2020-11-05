package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/MuShare/mail-sender-pool/config"
	"github.com/slack-go/slack"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger      *log.Logger
	slackClient *slack.Client
	logPrefix   = ""
	levelFlags  = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

//Setup xxx
func Setup() {
	if config.LogConfiguration.LogFilePath != "" {
		logger = log.New(&lumberjack.Logger{
			Filename:   config.LogConfiguration.LogFilePath,
			MaxSize:    500,
			MaxBackups: 2,
			MaxAge:     120,
			Compress:   true,
		}, logPrefix, log.LstdFlags)
	} else {
		logger = log.New(os.Stderr, logPrefix, log.LstdFlags)
	}
	if config.LogConfiguration.SlackToken != "" {
		slackClient = slack.New(config.LogConfiguration.SlackToken)
	}
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
	go slackSendMessage(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
	go slackSendMessage(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func slackSendMessage(v ...interface{}) {
	if slackClient != nil {
		slackClient.SendMessage(
			"easyjapanese-error",
			slack.MsgOptionUsername("mail pool test"),
			slack.MsgOptionIconEmoji(":anger:"),
			slack.MsgOptionText(fmt.Sprintf("%s   %v", logPrefix, v), true),
		)
	}
}
