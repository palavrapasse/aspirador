package pkg

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	datePattern     = "%d%" // the date in the local time zone: 2009/01/23
	fileNamePattern = "%f%" // full file name: /a/b/c/d.go
	levelPattern    = "%L%" // Level: TRACE
	linePattern     = "%l%" // line number: 23
	methodPattern   = "%M%" // method name: main
	messagePattern  = "%m%" // message to be logged: logging message
	timePattern     = "%t%" // the time in the local time zone: 01:23:23
)

const callDepth = 4

var Pattern = map[string]replaceFunction{
	datePattern:     applyDatePattern(),
	fileNamePattern: applyFileNamePattern(),
	levelPattern:    applyLevelPattern(),
	linePattern:     applyLinePattern(),
	methodPattern:   applyMethodPattern(),
	messagePattern:  applyMessagePattern(),
	timePattern:     applyTimePattern(),
}

type replaceFunction func(format string, r Record) string

func applyDatePattern() replaceFunction {
	return func(format string, r Record) string {
		year, month, day := time.Now().Date()
		return strings.Replace(format, datePattern, fmt.Sprintf("%04d/%02d/%02d", year, month, day), -1)
	}
}

func applyFileNamePattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		_, file, _, ok := runtime.Caller(callDepth)

		if ok {
			replaceWith = file
		}

		return strings.Replace(format, fileNamePattern, replaceWith, -1)
	}
}

func applyLevelPattern() replaceFunction {
	return func(format string, r Record) string {
		return strings.Replace(format, levelPattern, r.Level.String(), -1)
	}
}

func applyLinePattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		_, _, no, ok := runtime.Caller(callDepth)

		if ok {
			replaceWith = strconv.Itoa(no)
		}

		return strings.Replace(format, linePattern, replaceWith, -1)
	}
}

func applyMethodPattern() replaceFunction {
	return func(format string, r Record) string {
		replaceWith := ""
		pc, _, _, ok := runtime.Caller(callDepth)
		details := runtime.FuncForPC(pc)

		if ok && details != nil {
			replaceWith = details.Name()
		}

		return strings.Replace(format, methodPattern, replaceWith, -1)
	}
}

func applyMessagePattern() replaceFunction {
	return func(format string, r Record) string {
		return strings.Replace(format, messagePattern, r.Message, -1)
	}
}

func applyTimePattern() replaceFunction {
	return func(format string, r Record) string {
		hour, min, sec := time.Now().Clock()
		return strings.Replace(format, timePattern, fmt.Sprintf("%02d:%02d:%02d", hour, min, sec), -1)
	}
}
