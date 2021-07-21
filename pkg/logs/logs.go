package logs

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger = logrus.New()

func init() {

	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)
	logger.SetOutput(&lumberjack.Logger{
		Filename: "/logs/mlog.log",
		MaxSize:  10, // megabytes
		// MaxBackups: 3,
		// MaxAge:   28,   //days
		LocalTime: true,
		Compress:  true, // disabled by default
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true) // 建议只在 dev 环境下开启

	log.SetOutput(logger.Writer())
}

func Trace(args ...interface{}) {
	logger.Trace(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}
