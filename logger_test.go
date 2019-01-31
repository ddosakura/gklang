package gklang

import (
	"log"
	"os"
	"testing"
)

func init() {
	// l := log.New(os.Stdout, "[xxxx]: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	l := log.New(os.Stdout, "[xxxx]: ", log.LstdFlags)
	LoadLogger(l, LvDebug)
}

/*
func TestCrashLogger(t *testing.T) {
	defer func() {
		e := recover()
		println(e, "get it!\n")
	}()
	Log(LCrash, "Hello World!")
	println("ok!\n")
}
*/

func TestErrorLogger(t *testing.T) {
	defer func() {
		e := recover()
		println(e, "get it!\n")
	}()
	Log(LError, "Hello World!")
	println("ok!\n")
}

func TestWarnLogger(t *testing.T) {
	defer func() {
		e := recover()
		println(e, "get it!\n")
	}()
	Log(LWarn, "Hello World!")
	println("ok!\n")
}

func TestInfoLogger(t *testing.T) {
	defer func() {
		e := recover()
		println(e, "get it!\n")
	}()
	Log(LInfo, "Hello World!")
	println("ok!\n")
}

func TestDebugLogger(t *testing.T) {
	defer func() {
		e := recover()
		println(e, "get it!\n")
	}()
	Log(LDebug, "Hello World!", t)
	println("ok!\n")
}
