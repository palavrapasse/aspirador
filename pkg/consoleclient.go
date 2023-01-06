package pkg

import (
	"log"
	"os"
)

type ConsoleClient struct {
	loggers []*log.Logger
}

func NewConsoleClient() ConsoleClient {
	loggers := make([]*log.Logger, len(levelPrefix))

	for i, v := range levelPrefix {
		loggers[i] = log.New(os.Stdout, v, defaultLoggerFlag)
	}

	return ConsoleClient{
		loggers: loggers,
	}
}

func (cc ConsoleClient) Write(ar Record) {
	cc.loggers[ar.Level].Println(ar.Message)
}
