package log

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
)

func TraceId(ctx context.Context) zap.Field {
	return zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String())
}

func DefaultZapCore(fileName string, dir string, openLevel zapcore.Level) zapcore.Core {
	errEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level >= zap.ErrorLevel && zap.ErrorLevel >= openLevel
		})
	infoEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level == zap.InfoLevel && level >= openLevel
		})
	warnEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level == zap.WarnLevel && level >= openLevel
		})
	debugEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level == zap.DebugLevel && level >= openLevel
		})

	infoLogWriter := getLogWriter(path.Join(dir, fileName+"_info.log"))
	errLogWriter := getLogWriter(path.Join(dir, fileName+"_err.log"))
	warnWriter := getLogWriter(path.Join(dir, fileName+"_warn.log"))
	debugWriter := getLogWriter(path.Join(dir, fileName+"_debug.log"))

	stdout := zapcore.AddSync(os.Stdout)
	infoCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(infoLogWriter, stdout), infoEnable)
	errCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(errLogWriter, stdout), errEnable)
	warnCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(warnWriter, stdout), warnEnable)
	debugCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(debugWriter, stdout), debugEnable)

	return zapcore.NewTee(infoCore, errCore, warnCore, debugCore)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
