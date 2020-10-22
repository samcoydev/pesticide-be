package logHandler

import (
	log "github.com/sirupsen/logrus"
)

func Debug(msg string) {
	log.Debug(msg)
}

func Err(msg string) {
	log.Error(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}
