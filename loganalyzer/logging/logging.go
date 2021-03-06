package logging

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger.Level = logrus.DebugLevel
}

func SetLevel(level string) {
	switch strings.ToLower(level) {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
}

func SetOutputFile(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logger.Fatalln("Open log file error:", err)
	}
	logger.Out = logFile
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Info(args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}
