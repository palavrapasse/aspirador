package pkg

import (
	"log"
)

const (
	defaultLoggerFlag = log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds
)

type Aspirador struct {
	writers []AspiradorClient
}

// Default to Console.
func NewAspirador() Aspirador {
	writers := make([]AspiradorClient, 1)
	writers[0] = NewConsoleClient()

	return Aspirador{
		writers: writers,
	}
}

func (as *Aspirador) AddAspirador(fp string) {
	writer := NewFileClient(fp)

	as.writers = append(as.writers, writer)
}

func (as Aspirador) Trace(msg string) {
	as.log(TRACE, msg)
}

func (as Aspirador) Info(msg string) {
	as.log(INFO, msg)
}

func (as Aspirador) Warning(msg string) {
	as.log(WARNING, msg)
}

func (as Aspirador) Error(msg string) {
	as.log(ERROR, msg)
}

func (as Aspirador) log(lvl Level, msg string) {
	record := AspiradorRecord{
		Level:   lvl,
		Message: msg,
	}

	for _, v := range as.writers {
		v.Write(record)
	}

}
