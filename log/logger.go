package log

import (
	"github.com/sirupsen/logrus"
)

var standardLogger = logrus.StandardLogger()

func init() {
	standardLogger.SetLevel(logrus.DebugLevel)
}

func Debug(msg string) {
	Debugf(msg)
}

func Debugf(format string, args ...interface{}) {
	standardLogger.Debugf(format, args)
}

func Errorf(format string, args ...interface{}) {
	standardLogger.Errorf(format, args)
}
