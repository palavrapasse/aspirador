package pkg

import "fmt"

var defaultPatternLayout = PatternLayout(fmt.Sprintf("[%s] %s %s %s.%s:%s : %s", LevelPattern, DatePattern, TimePattern, FileNamePattern, MethodPattern, LinePattern, MessagePattern))

type PatternLayout string

func (p PatternLayout) FormatRecord(r Record) string {
	if len(p) == 0 {
		return ""
	}

	result := string(p)
	for _, p := range pattern {
		result = p(result, r)
	}

	return result
}
