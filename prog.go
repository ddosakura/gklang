package gklang

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// DevMode define the default logger level
	// Enable In Linux:
	//     export SAKURA_GK_DEV=true
	DevMode = false
)

// Init App
func Init(name string, envFile ...string) {
	// Init Env
	if len(envFile) > 0 {
		err := godotenv.Load(envFile...)
		if err != nil {
			Er(err)
		}
	}

	// Init Log
	l := log.New(os.Stdout, "["+name+"]: ", log.LstdFlags)
	if os.Getenv("SAKURA_GK_DEV") == "true" {
		DevMode = true
	}
	if DevMode {
		LoadLogger(l, LvDebug)
	} else {
		LoadLogger(l, LvInfo)
	}
}

// Finish with Errs and Warns
func Finish(code int, v ...string) {
	if len(v) > 0 {
		Log(LCrash, fmt.Sprintf(v[0], Errs, Warns))
	}
	os.Exit(code)
}
