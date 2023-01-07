package pkg

import (
	"log"
	"os"
)

type FileClient struct {
	loggers []*log.Logger
}

func NewFileClient(fp string) (FileClient, error) {
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)

	if err != nil {
		return FileClient{}, err
	}

	loggers := make([]*log.Logger, len(levelPrefix))

	for i, v := range levelPrefix {
		loggers[i] = log.New(file, v, defaultLoggerFlag)
	}

	return FileClient{
		loggers: loggers,
	}, nil
}

func (fc FileClient) Write(ar Record) {
	fc.loggers[ar.Level].Println(ar.Message)
}
