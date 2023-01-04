package pkg

import (
	"io"
	"log"
	"os"
)

var stdout io.Writer = os.Stdout

type ConsoleWriter struct {
	logs []*log.Logger
}

func NewConsoleWriter() ConsoleWriter {
	logs := make([]*log.Logger, len(levelStrings))

	for i, v := range levelStrings {
		logs[i] = log.New(stdout, v, defaultLoggerFlag)
	}

	return ConsoleWriter{
		logs: logs,
	}
}

func (cw ConsoleWriter) Write(ar aspiradorRecord) {
	log := cw.logs[ar.Level]

	if log == nil {
		//ERROR
		return
	}

	log.Println(ar.Message)
}
