package gklang

import (
	"fmt"
	"os"
	"testing"
)

const (
	EnvText     = "DS_GK_LANG_LOG_LEVEL"
	EnvLogLevel = "31"

	ErrText = "finish with %d error(s) and %d warning(s)"
)

func TestInit(t *testing.T) {
	Init("test")
	Log(LInfo, "init success")
	lv := os.Getenv(EnvText)
	if lv != EnvLogLevel {
		t.Fatalf("The DS_GK_LANG_MODE should be `%s`", EnvLogLevel)
	}
	fmt.Printf(ErrText, Errs, Warns)
}

func TestFinish(t *testing.T) {
	Finish(0, ErrText)
}
