package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	std = zap.New(zapcore.NewCore(getEncoder(), os.Stdout, zap.DebugLevel), zap.AddCaller())
)

func Info(msg string, fields ...zapcore.Field) {
	std.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	std.Debug(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	std.Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	std.Fatal(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	std.Warn(msg, fields...)
}
func Panic(msg string, fields ...zapcore.Field) {
	std.Panic(msg, fields...)
}

func Sync() error {
	return std.Sync()
}
