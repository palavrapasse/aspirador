package pkg

import (
	"os"
)

type FileClient struct {
	loggers       LevelLogger
	patternLayout PatternLayout
}

func NewFileClient(fp string, levels ...Level) (FileClient, error) {
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)

	if err != nil {
		return FileClient{}, err
	}

	loggers := NewLevelLogger(file, defaultLoggerFlag, levels)

	return FileClient{
		loggers:       loggers,
		patternLayout: defaultPatternLayout,
	}, nil
}

func (fc *FileClient) SetPatternLayout(p PatternLayout) {
	fc.patternLayout = p
}

func (fc FileClient) Write(ar Record) {
	logger, exists := fc.loggers[ar.Level]

	if !exists {
		return
	}

	message := fc.patternLayout.FormatRecord(ar)

	logger.Println(message)
}

func (fc FileClient) SupportsLevel(l Level) bool {
	return fc.loggers.ContainsLevel(l)
}
