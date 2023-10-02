package logger

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}

func Init() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "server.log",
		MaxSize:    1, // 单一档案最大几M
		MaxBackups: 3, // 最多保留几份
		MaxAge:     1, // 最多保留几天
		Compress:   true,
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:    "level",
		TimeKey:     "time",
		MessageKey:  "message",
		NameKey:     "name", // 可以放自定义x-api-id
		CallerKey:   "caller",
		FunctionKey: "func",
		// StacktraceKey: "trace",
		// LineEnding:     "\r\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// EncodeCaller:   MyCaller, // 自定义
	})
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller(),
		zap.AddCallerSkip(1),
		// zap.AddStacktrace(zap.DebugLevel),
	)
}

func Close() {
	logger.Sync()
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func NameInfof(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Infof(template, args...)
}

func NameErrorf(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Errorf(template, args...)
}

func NameWarnf(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Warnf(template, args...)
}

func NameDebugf(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Debugf(template, args...)
}
