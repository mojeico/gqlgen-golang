package logger

import (
	"context"
	"github.com/mojeico/gqlgen-golang/internal/middleware"
	"github.com/sirupsen/logrus"
	"time"
)

func CrudLoggerError(ctx context.Context, method, file, errorLogMessage string) {

	currentUser, _ := middleware.GetCurrentUserFromContext(ctx)

	logrus.WithFields(logrus.Fields{
		"method":   method,
		"file":     file,
		"username": currentUser.Username,
		"email":    currentUser.Email,
		"time":     time.Now().Format("01-02-2006 15:04:05"),
	}).Error(errorLogMessage)

}

func CrudLoggerInfo(ctx context.Context, method, file, errorLogMessage string) {

	currentUser, _ := middleware.GetCurrentUserFromContext(ctx)

	logrus.WithFields(logrus.Fields{
		"method":   method,
		"file":     file,
		"username": currentUser.Username,
		"email":    currentUser.Email,
		"time":     time.Now().Format("01-02-2006 15:04:05"),
	}).Info(errorLogMessage)
}
