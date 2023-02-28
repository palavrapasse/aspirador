package pkg

import (
	"os"
)

type ConsoleClient struct {
	loggers       LevelLogger
	patternLayout PatternLayout
}

func NewConsoleClient(levels ...Level) ConsoleClient {
	loggers := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)

	return ConsoleClient{
		loggers:       loggers,
		patternLayout: defaultPatternLayout,
	}
}

func (cc *ConsoleClient) SetPatternLayout(p PatternLayout) {
	cc.patternLayout = p
}

func (cc ConsoleClient) Write(ar Record) {
	logger, exists := cc.loggers[ar.Level]

	if !exists {
		return
	}

	message := cc.patternLayout.FormatRecord(ar)

	logger.Println(message)
}

func (cc ConsoleClient) SupportsLevel(l Level) bool {
	return cc.loggers.ContainsLevel(l)
}
