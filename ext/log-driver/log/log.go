package log

import (
	"sync"

	lg "github.com/sirupsen/logrus"
)

var _logger *lg.Logger
var _loggerOnce sync.Once

// init init
func init() {
	_ = NewLogger()
	NewLogger().SetFormatter(&lg.JSONFormatter{})
}

// NewLogger logger
func NewLogger() *lg.Logger {
	if _logger != nil {
		return _logger
	}
	_loggerOnce.Do(func() {
		_logger = lg.New()
		_logger.WithFields(lg.Fields{
			"project": "t-openTrace",
		})
	})
	return _logger
}

// Warn warn
func Warn(args ...interface{}) {
	NewLogger().Warnln(args...)
}

// Info info
func Info(args ...interface{}) {
	NewLogger().Infoln(args...)
}

// Debug dev
func Debug(args ...interface{}) {
	NewLogger().Debugln(args...)
}

// DebugF dev
func DebugF(f string, args ...interface{}) {
	NewLogger().Debugf(f, args...)
}

// Trace dev
func Trace(args ...interface{}) {
	NewLogger().Traceln(args...)
}

// Fatal dev
func Fatal(args ...interface{}) {
	NewLogger().Fatalln(args...)
}

// Error dev
func Error(args ...interface{}) {
	NewLogger().Errorln(args...)
}
