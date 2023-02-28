package pkg

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	DatePattern     = "%d%" // the date in the local time zone: 2009/01/23
	FileNamePattern = "%f%" // full file name: /a/b/c/d.go
	LevelPattern    = "%L%" // Level: TRACE
	LinePattern     = "%l%" // line number: 23
	MethodPattern   = "%M%" // method name: main
	MessagePattern  = "%m%" // message to be logged: logging message
	TimePattern     = "%t%" // the time in the local time zone: 01:23:23
)

const callDepth = 5

var pattern = map[string]replaceFunction{
	DatePattern:     applyDatePattern(),
	FileNamePattern: applyFileNamePattern(),
	LevelPattern:    applyLevelPattern(),
	LinePattern:     applyLinePattern(),
	MethodPattern:   applyMethodPattern(),
	MessagePattern:  applyMessagePattern(),
	TimePattern:     applyTimePattern(),
}

type replaceFunction func(format string, r Record) string

func applyDatePattern() replaceFunction {
	return func(format string, r Record) string {
		year, month, day := time.Now().Date()
		return strings.Replace(format, DatePattern, fmt.Sprintf("%04d/%02d/%02d", year, month, day), -1)
	}
}

func applyFileNamePattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		_, file, _, ok := runtime.Caller(callDepth)

		if ok {
			replaceWith = getLastStringOfSeparator(file, '/', 2)
		}

		return strings.Replace(format, FileNamePattern, replaceWith, -1)
	}
}

func applyLevelPattern() replaceFunction {
	return func(format string, r Record) string {
		return strings.Replace(format, LevelPattern, r.Level.String(), -1)
	}
}

func applyLinePattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		_, _, no, ok := runtime.Caller(callDepth)

		if ok {
			replaceWith = strconv.Itoa(no)
		}

		return strings.Replace(format, LinePattern, replaceWith, -1)
	}
}

func applyMethodPattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		pc, _, _, ok := runtime.Caller(callDepth)
		details := runtime.FuncForPC(pc)

		if ok && details != nil {
			name := details.Name()
			replaceWith = getLastStringOfSeparator(name, '.', 1)
		}

		return strings.Replace(format, MethodPattern, replaceWith, -1)
	}
}

func applyMessagePattern() replaceFunction {
	return func(format string, r Record) string {
		return strings.Replace(format, MessagePattern, r.Message, -1)
	}
}

func applyTimePattern() replaceFunction {
	return func(format string, r Record) string {
		hour, min, sec := time.Now().Clock()
		return strings.Replace(format, TimePattern, fmt.Sprintf("%02d:%02d:%02d", hour, min, sec), -1)
	}
}

func getLastStringOfSeparator(value string, separator byte, count int) string {
	short := value
	c := 0

	for i := len(value) - 1; i > 0; i-- {
		if value[i] == separator {
			c++

			if c == count {
				short = value[i+1:]
				break
			}
		}
	}

	return short
}
