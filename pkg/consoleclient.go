package pkg

import (
	"os"
)

type ConsoleClient struct {
	loggers LevelLogger
}

func NewConsoleClient(levels ...Level) ConsoleClient {
	loggers := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)

	return ConsoleClient{
		loggers: loggers,
	}
}

func (cc ConsoleClient) Write(ar Record) {
	cc.loggers[ar.Level].Println(ar.Message)
}
