package pkg

import (
	"io"
	"log"
)

type LevelLogger map[Level]*log.Logger

func NewLevelLogger(out io.Writer, flag int, levels []Level) LevelLogger {
	size := len(levels)

	if size == 0 {
		levels = []Level{TRACE, INFO, WARNING, ERROR}
		size = len(levels)
	}

	result := make(map[Level]*log.Logger, size)

	for _, v := range levels {
		result[v] = log.New(out, levelPrefix[v], flag)
	}

	return result
}

func (ll LevelLogger) ContainsLevel(l Level) bool {
	_, exists := ll[l]
	return exists
}
