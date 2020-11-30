package log

import (
	"github.com/sirupsen/logrus"
)

var standardLogger = logrus.StandardLogger()

func Debugf(format string, args ...interface{}) {
	standardLogger.Debugf(format, args)
}

func Errorf(format string, args ...interface{}) {
	standardLogger.Errorf(format, args)
}
