package log

import (
	"github.com/sirupsen/logrus"
)

var standardLogger = logrus.StandardLogger()

func init() {
	standardLogger.SetLevel(logrus.DebugLevel)
	standardLogger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
	})
}

func Debug(msg string) {
	standardLogger.Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	standardLogger.Debugf(format, args)
}

func Errorf(format string, args ...interface{}) {
	standardLogger.Errorf(format, args)
}
