package pkg

import (
	"log"
)

const (
	defaultLoggerFlag = log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds
)

type Aspirador struct {
	clients []Client
}

// Default to Console.
func NewAspirador() Aspirador {
	clients := make([]Client, 1)
	clients[0] = NewConsoleClient()

	return Aspirador{
		clients: clients,
	}
}

func (as *Aspirador) AddFileClient(fp string) error {
	client, err := NewFileClient(fp)

	if err == nil {
		as.clients = append(as.clients, client)
	}

	return err
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
	record := Record{
		Level:   lvl,
		Message: msg,
	}

	for _, v := range as.clients {
		v.Write(record)
	}

}
