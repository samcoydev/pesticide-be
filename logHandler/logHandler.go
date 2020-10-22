package logHandler

import (
	"github.com/sirupsen/logrus"
	"time"
)

var log = logrus.New()

func InitLog(fromName string, msg string) {
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "]"

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		DisableQuote:     true,
		ForceColors:      true,
	})
	log.Info(timeStamp, fromName+" ", msg)
}

func Debug(fromName string, msg string) {
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "]"

	log.Debug(timeStamp, fromName+" ", msg)
}

func Info(fromName string, msg string) {
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "]"

	log.Info(timeStamp, fromName+" ", msg)
}

func Err(fromName string, msg string) {
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "]"

	log.Error(timeStamp, fromName+" ", msg)
}

func Warn(fromName string, msg string) {
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "]"

	log.Warn(timeStamp, fromName+" ", msg)
}
