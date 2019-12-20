package jotnar

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Logger

func InitLogrus() {
	defaultLogger = logrus.New()
	switch strings.ToLower(defualtLogConfig.Level) {
	case "panic":
		defaultLogger.Level = logrus.PanicLevel
	case "fatal":
		defaultLogger.Level = logrus.FatalLevel
	case "error":
		defaultLogger.Level = logrus.ErrorLevel
	case "Warn":
		defaultLogger.Level = logrus.WarnLevel
	case "info":
		defaultLogger.Level = logrus.InfoLevel
	case "debug":
		defaultLogger.Level = logrus.DebugLevel
	case "trace":
		defaultLogger.Level = logrus.TraceLevel
	}
}

