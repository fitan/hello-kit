package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type option struct {
	traceLevel zapcore.Level
	openTrace  bool
	logLevel   zapcore.Level
	dir        string
	fileName   string
	filter     map[string]struct{}
	*zap.Logger
}

type Option func(o *option)

func WithLogger(logger *zap.Logger) Option {
	return func(o *option) {
		o.Logger = logger
	}
}

func WithLogFileName(fileName string, dir string) Option {
	return func(o *option) {
		o.fileName = fileName
		o.dir = dir
	}
}

func WithLogLevel(level zapcore.Level) Option {
	return func(o *option) {
		o.logLevel = level
	}
}

func WithTrace(level zapcore.Level, filter map[string]struct{}) Option {
	return func(o *option) {
		o.traceLevel = level
		//o.tp = tp
		o.openTrace = true

		o.filter = filter

	}
}
