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

func (cw ConsoleClient) Write(ar AspiradorRecord) {
	logger := cw.loggers[ar.Level]
	logger.Println(ar.Message)
}
