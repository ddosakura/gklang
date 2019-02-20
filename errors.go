package gklang

import (
	"log"
)

// Er exit with error log
func Er(e error) {
	Log(LCrash, e)
}

// Err used when no logger has be loaded
func Err(e error) {
	log.Fatalln(e)
}
