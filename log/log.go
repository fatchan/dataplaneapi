package log

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	apache_log "github.com/lestrrat-go/apache-logformat"
	"github.com/sirupsen/logrus"
)

const (
	PanicLevel logrus.Level = logrus.PanicLevel
	FatalLevel logrus.Level = logrus.FatalLevel
	ErrorLevel logrus.Level = logrus.ErrorLevel
	WarnLevel  logrus.Level = logrus.WarnLevel
	InfoLevel  logrus.Level = logrus.InfoLevel
	DebugLevel logrus.Level = logrus.DebugLevel
	TraceLevel logrus.Level = logrus.TraceLevel
)

var (
	appLogger       *Logger
	accessLogger    *Logger
	config          Targets
	oldConfig       LoggingOptions
	oldSyslogConfig SyslogOptions
)

func InitWithConfiguration(targets Targets, old LoggingOptions, oldSyslog SyslogOptions) error {
	config = targets
	oldConfig = old
	oldSyslogConfig = oldSyslog
	return Init()
}

func Init() error {
	if len(config) > 0 {
		for _, target := range config {
			switch len(target.LogTypes) {
			case 0:
				// by default use only app loggers
				configureAppLogger(target)
			case 1:
				switch target.LogTypes[0] {
				case "app":
					configureAppLogger(target)
				case "access":
					configureAccessLogger(target)
				}
			case 2:
				if target.LogTypes[0] != target.LogTypes[1] {
					configureAppLogger(target)
					configureAccessLogger(target)
				}
			}
		}
	} else {
		// Deprecated: if no log targets are set in the configuration file, use old way
		target := Target{
			LogTo:          oldConfig.LogTo,
			LogLevel:       oldConfig.LogLevel,
			LogFile:        oldConfig.LogFile,
			LogFormat:      oldConfig.LogFormat,
			LogTypes:       []string{"access", "app"},
			ACLFormat:      oldConfig.ACLFormat,
			SyslogAddr:     oldSyslogConfig.SyslogAddr,
			SyslogProto:    oldSyslogConfig.SyslogProto,
			SyslogTag:      oldSyslogConfig.SyslogTag,
			SyslogLevel:    oldSyslogConfig.SyslogLevel,
			SyslogFacility: oldSyslogConfig.SyslogFacility,
		}
		configureAccessLogger(target)
		configureAppLogger(target)
	}
	return nil
}

func configureAppLogger(target Target) {
	logger := logrus.New()
	configureLogger(logger, target)
	if appLogger == nil {
		appLogger = &Logger{}
	}
	aclLogger := &ACLLogger{
		Logger: logger,
	}
	appLogger.loggers = append(appLogger.loggers, aclLogger)
}

func configureAccessLogger(target Target) {
	logger := logrus.New()
	configureLogger(logger, target)
	if accessLogger == nil {
		accessLogger = &Logger{}
	}
	var al *apache_log.ApacheLog
	if target.ACLFormat == "" {
		al, _ = apache_log.New(DefaultApacheLogFormat)
	} else {
		al, _ = apache_log.New(target.ACLFormat)
	}
	if al == nil {
		return
	}
	aclLogger := &ACLLogger{
		Logger:    logger,
		ApacheLog: al,
	}
	accessLogger.loggers = append(accessLogger.loggers, aclLogger)
}

func AppLogger() (*Logger, error) {
	if appLogger == nil {
		if err := Init(); err != nil {
			return nil, err
		}
	}
	if appLogger == nil {
		return nil, errors.New("no application loggers configured")
	}
	return appLogger, nil
}

func AccessLogger() (*Logger, error) {
	if accessLogger == nil {
		if err := Init(); err != nil {
			return nil, err
		}
	}
	if accessLogger == nil {
		return nil, errors.New("no access loggers configured")
	}
	return accessLogger, nil
}

func Log(level logrus.Level, args ...interface{}) {
	if appLogger != nil {
		appLogger.Log(level, args...)
	}
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Logf(level, format, args...)
	}
}

func Print(args ...interface{}) {
	if appLogger != nil {
		appLogger.Print(args...)
	}
}

func Trace(args ...interface{}) {
	if appLogger != nil {
		appLogger.Trace(args...)
	}
}

func Debug(args ...interface{}) {
	if appLogger != nil {
		appLogger.Debug(args...)
	}
}

func Info(args ...interface{}) {
	if appLogger != nil {
		appLogger.Info(args...)
	}
}

func Warning(args ...interface{}) {
	if appLogger != nil {
		appLogger.Warning(args...)
	}
}

func Error(args ...interface{}) {
	if appLogger != nil {
		appLogger.Error(args...)
	}
}

func Panic(args ...interface{}) {
	if appLogger != nil {
		appLogger.Panic(args...)
	}
}

func Fatal(args ...interface{}) {
	if appLogger != nil {
		appLogger.Panic(args...)
	}
}

func Printf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Printf(format, args...)
	}
}

func Tracef(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Tracef(format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Debugf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Infof(format, args...)
	}
}

func Warningf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Warningf(format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Errorf(format, args...)
	}
}

func Panicf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Panicf(format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Fatalf(format, args...)
	}
}

func Fatalln(format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Fatalln(args...)
	}
}

func WithFields(fields map[string]interface{}, level logrus.Level, args ...interface{}) {
	if appLogger != nil {
		appLogger.WithFields(fields, level, args...)
	}
}

func WithFieldsf(fields map[string]interface{}, level logrus.Level, format string, args ...interface{}) {
	if appLogger != nil {
		appLogger.WithFieldsf(fields, level, format, args...)
	}
}

func configureLogger(logger *logrus.Logger, target Target) {
	switch target.LogFormat {
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			DisableColors: true,
		})
	case "JSON":
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	switch target.LogTo {
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "file":
		dir := filepath.Dir(target.LogFile)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			Warning("Error opening log file, no logging implemented: " + err.Error())
		}
		//nolint:govet
		logFile, err := os.OpenFile(target.LogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			Warning("Error opening log file, no logging implemented: " + err.Error())
		}
		logger.SetOutput(logFile)
	case "syslog":
		logger.SetOutput(ioutil.Discard)
		hook, err := NewRFC5424Hook(target)
		if err != nil {
			Warningf("Error configuring Syslog logging: %s", err.Error())
			break
		}
		logger.AddHook(hook)
	}

	switch target.LogLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	}
}