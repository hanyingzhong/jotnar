package jotnar

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Logger

func initLogrus() {
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

	if defualtLogConfig.Format == "json" {
		jsonFormat := &logrus.JSONFormatter{
			CallerPrettyfier: myCallerPrettyfier,
			PrettyPrint:      defualtLogConfig.IsPretty,
			TimestampFormat:  defualtLogConfig.Timeformat,
		}

		defaultLogger.SetFormatter(jsonFormat)
	} else if defualtLogConfig.Format == "text" {
		textFormat := defaultLogger.Formatter.(*logrus.TextFormatter)
		textFormat.CallerPrettyfier = myCallerPrettyfier
		textFormat.TimestampFormat = defualtLogConfig.Timeformat
		textFormat.FullTimestamp = true
	} else {
		errExit(errors.New("format value must be json or text"))
	}

	defaultLogger.SetReportCaller(true)

	fileName := defualtLogConfig.FilePath
	if fileName != "" {
		var (
			f   *os.File
			err error
		)

		if _, err = os.Stat(fileName); err != nil {
			if os.IsNotExist(err) {
				f, err = os.Create(fileName)
				errExit(err)
			} else {
				errExit(err)
			}
		} else {
			f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			errExit(err)
		}

		defaultLogger.SetOutput(io.MultiWriter(f, os.Stdout))
	}
}

func myCallerPrettyfier(f *runtime.Frame) (string, string) {
	s := strings.Split(f.Function, ".")
	funcname := s[len(s)-1]
	dir, filename := path.Split(f.File)
	tmpArray := strings.Split(dir, string(os.PathSeparator))
	if len(tmpArray) <= 3 {
		filename = f.File
	} else {
		tmpArray = tmpArray[len(tmpArray)-3:]
		filename = strings.Join(tmpArray, string(os.PathSeparator)) + filename
	}
	return funcname, filename + ":" + fmt.Sprint(f.Line)
}

func GetLogger() *logrus.Logger {
	return defaultLogger
}
