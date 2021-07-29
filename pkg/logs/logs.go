package logs

import (
	"log"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.InfoLevel)
	// logger.SetOutput(os.Stdout)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "/logs/mlog.log",
		MaxSize:    10,
		MaxAge:     0,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   true,
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetReportCaller(true) // 建议只在 dev 环境下开启

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

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

func Traceln(args ...interface{}) {
	logger.Traceln(args)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args)
}
