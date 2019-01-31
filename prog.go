package gklang

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Init App
func Init(name string) {
	// Init Log
	l := log.New(os.Stdout, "["+name+"]: ", log.LstdFlags)
	LoadLogger(l, LvDebug)

	// Init Env
	err := godotenv.Load()
	if err != nil {
		Er(err)
	}
}

// Finish with Errs and Warns
func Finish(code int, v ...string) {
	if len(v) > 0 {
		Log(LCrash, fmt.Sprintf(v[0], Errs, Warns))
	}
	os.Exit(code)
}
