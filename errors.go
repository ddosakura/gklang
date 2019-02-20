package gklang

import (
	"os"
)

// Er exit with error log
func Er(e error) {
	Log(LCrash, e)
	os.Exit(1)
}
