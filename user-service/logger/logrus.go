package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logFile *os.File

func InitLog() {
	var err error
	logFile, err = os.OpenFile("logs/user-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
}

func Debug(msg ...interface{}) {
	logrus.Debug(msg...)
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func Fatal(msg ...interface{}) {
	logrus.Fatal(msg...)
}

