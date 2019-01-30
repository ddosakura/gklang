package gklang

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

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
)

// LogLevel is the level of log
type LogLevel int

// LogHandler can define custom log level
type LogHandler func(level LogLevel, v ...interface{})

var (
	loggers = make(map[*log.Logger]LogLevel)
	// ExLevelHandler can use custom log level
	ExLevelHandler LogHandler = nil
)

// LoadLogger save logger to loggers
func LoadLogger(l *log.Logger, lv LogLevel) {
	loggers[l] = lv
}

// Log message with level
func Log(level LogLevel, v ...interface{}) {
	for l, lv := range loggers {
		LV := lv & level
		if LV == 0 {
			continue
		}
		if LV&LCrash != 0 {
			l.Fatalln(v...)
			continue
		}
		if LV&LError != 0 {
			l.Panicln(v...)
			continue
		}
		if LV&LWarn != 0 {
			l.Println(v...)
			continue
		}
		if LV&LInfo != 0 {
			l.Println(v...)
			continue
		}
		if LV&LDebug != 0 {
			spew.Println(v...)
			continue
		}
		if ExLevelHandler != nil {
			ExLevelHandler(level, v...)
		}
	}
}
