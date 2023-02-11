package pkg

import (
	"os"
)

type FileClient struct {
	loggers LevelLogger
}

func NewFileClient(fp string, levels ...Level) (FileClient, error) {
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)

	if err != nil {
		return FileClient{}, err
	}

	loggers := NewLevelLogger(file, defaultLoggerFlag, levels)

	return FileClient{
		loggers: loggers,
	}, nil
}

func (fc FileClient) Write(ar Record) {
	fc.loggers[ar.Level].Println(ar.Message)
}
