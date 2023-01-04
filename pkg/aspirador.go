package pkg

import (
	"log"
)

const (
	defaultLoggerFlag = log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds
)

var (
	global Aspirador
)

type aspiradorRecord struct {
	Message string
	Level   Level
}

type aspiradorWriter interface {
	Write(ar aspiradorRecord)
}

type Aspirador struct {
	writers []aspiradorWriter
}

// Default to Console.
func NewAspirador() {
	writers := make([]aspiradorWriter, 1)
	writers[0] = NewConsoleWriter()

	global = Aspirador{
		writers: writers,
	}
}

func Trace(msg string) {
	global.log(TRACE, msg)
}

func Info(msg string) {
	global.log(INFO, msg)
}

func Warning(msg string) {
	global.log(WARNING, msg)
}

func Error(msg string) {
	global.log(ERROR, msg)
}

func (as Aspirador) log(lvl Level, msg string) {
	record := aspiradorRecord{
		Level:   lvl,
		Message: msg,
	}

	for _, v := range as.writers {
		v.Write(record)
	}

}
