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
	logger, exists := cc.loggers[ar.Level]

	if !exists {
		return
	}

	logger.Println(ar.Message)
}
