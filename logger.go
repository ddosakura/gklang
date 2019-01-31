package gklang

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
)

// LogLevel is the level of log
type LogLevel int

// LogHandler can define custom log level
type LogHandler func(level LogLevel, v ...interface{})

const (
	// LCrash is a error message and exit program
	LCrash = 1 << iota
	// LError is a error message
	LError
	// LWarn is a warning message
	LWarn
	// LInfo is a normal message
	LInfo
	// LDebug is a debug message
	LDebug

	// LvCrash only log LCrach
	LvCrash = LCrash
	// LvError log level above LError
	LvError = LError | LvCrash
	// LvWarn log level above LWarn
	LvWarn = LWarn | LvError
	// LvInfo log level above LInfo
	LvInfo = LInfo | LvWarn
	// LvDebug all log
	LvDebug = LDebug | LvInfo

	colf = "%c[%d;;%dm%s%c[0m"
)

var (
	// tags
	cTag = ""
	eTag = fmt.Sprintf(colf, 0x1B, 1, 31, "[error]", 0x1B)
	wTag = fmt.Sprintf(colf, 0x1B, 1, 33, " [warn]", 0x1B)
	iTag = fmt.Sprintf(colf, 0x1B, 0, 32, " [info]", 0x1B)
	dTag = fmt.Sprintf(colf, 0x1B, 1, 34, "[debug]", 0x1B)

	loggers = make(map[*log.Logger]LogLevel)
	// ExLevelHandler can use custom log level
	ExLevelHandler LogHandler
	// Errs count the errors in program
	Errs = 0
	// Warns count the warnings in program
	Warns = 0
)

// LoadLogger save logger to loggers
func LoadLogger(l *log.Logger, lv LogLevel) {
	loggers[l] = lv
}

func buildMsg(tag string, v []interface{}) []interface{} {
	return append([]interface{}{tag}, v...)
}

// Log message with level
func Log(level LogLevel, v ...interface{}) {
	for l, lv := range loggers {
		LV := lv & level
		if LV == 0 {
			continue
		}
		if LV&LCrash != 0 {
			l.Fatalln(buildMsg(cTag, v)...)
			continue
		}
		if LV&LError != 0 {
			Errs++
			l.Panicln(buildMsg(eTag, v)...)
			continue
		}
		if LV&LWarn != 0 {
			Warns++
			l.Println(buildMsg(wTag, v)...)
			continue
		}
		if LV&LInfo != 0 {
			l.Println(buildMsg(iTag, v)...)
			continue
		}
		if LV&LDebug != 0 {
			spew.Println(buildMsg(dTag, v)...)
			continue
		}
		if ExLevelHandler != nil {
			ExLevelHandler(level, v...)
		}
	}
}
