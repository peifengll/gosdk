package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestDefauleLog(t *testing.T) {
	Info("info")
	Error("error")
	//Fatal("fatal")
	Warn("warn")
	//Panic("panic")
}

func TestNewLogger(t *testing.T) {
	SetStdout(true)
	logger := NewLogger(zap.Fields(zap.String("id", "test"), zap.String("reqpath", "/cxc/cxcx")))
	logger.Info("info")
	logger = NewLogger()
	logger.Info("info")

}
