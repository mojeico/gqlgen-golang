package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

func SystemLoggerError(method, file, errorLogMessage string) {

	logrus.WithFields(logrus.Fields{
		"method": method,
		"file":   file,
		"time":   time.Now().Format("01-02-2006 15:04:05"),
	}).Error(errorLogMessage)

}
