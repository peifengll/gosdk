package log

import "testing"

func TestDefauleLog(t *testing.T) {
	Info("info")
	Error("error")
	//Fatal("fatal")
	Warn("warn")
	//Panic("panic")
}
