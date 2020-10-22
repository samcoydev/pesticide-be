package logHandler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func Debug(from string, msg string) {
	fromName := "[" + from + "] "
	timeStamp := "[" + time.Now().Format("Jan _2 15:04:05") + "] "

	fmt.Println(fromName, timeStamp, msg)
}

func Err(msg string) {
	log.Error(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}
