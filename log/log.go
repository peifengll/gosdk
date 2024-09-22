package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	if stdout {
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(encoder, os.Stdout, level),
		)
		return zap.New(core, zap.AddCaller())
	} else {
		core := zapcore.NewCore(encoder, writeSyncer, level)
		return zap.New(core, zap.AddCaller())
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
