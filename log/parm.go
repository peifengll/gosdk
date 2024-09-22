package log

import (
	"go.uber.org/zap/zapcore"
)

var (
	filePath                 = "log.log"
	stdout                   = false
	level      zapcore.Level = -1
	maxSize                  = 500
	maxBackups               = 30
	maxAge                   = 30
	compress                 = false
)

func SetFilePath(newFilePath string) {
	filePath = newFilePath
}

func SetLevel(newLevel zapcore.Level) {
	level = newLevel
}

func SetMaxSize(newMaxSize int) {
	maxSize = newMaxSize
}

func SetMaxBackups(newMaxBackups int) {
	maxBackups = newMaxBackups
}

func SetMaxAge(newMaxAge int) {
	maxAge = newMaxAge
}

func SetCompress(newCompress bool) {
	compress = newCompress
}

func SetStdout(flag bool) {
	stdout = flag
}
