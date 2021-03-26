package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Logger is a configured logrus.Logger.
	Logger *logrus.Logger
	// Entry is a configured logrus.Entry.
	Entry *logrus.Entry
)

// StructuredLogger is to stores loggers configs
type StructuredLogger struct {
	Logger *logrus.Logger
	Entry  *logrus.Entry
}

// NewLogger creates and configures a new logrus Logger.
func NewLogger(module string) *StructuredLogger {
	// Setting up logger level and format
	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{}
	level := viper.GetString("log-level")
	if level == "" {
		level = "debug"
	}
	l, err := logrus.ParseLevel(level)
	if err != nil {
		level = "info"
	}
	Logger.Level = l

	// Setting up default fields
	logFields := logrus.Fields{}
	logFields["application"] = "pomoday-backend"
	logFields["module"] = module
	logFields["environment"] = viper.GetString("env")
	logFields["version"] = "1.0.0" //FIXME
	Entry = Logger.WithFields(logFields)

	// Saving logger config
	sl := &StructuredLogger{
		Logger: Logger,
		Entry:  Entry,
	}
	return sl
}
