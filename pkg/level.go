package pkg

const (
	TRACE = iota
	INFO
	WARNING
	ERROR
)

var (
	levelStrings = [...]string{"TRACE: ", "INFO: ", "WARNING: ", "ERROR: "}
)

type Level int

func (l Level) String() string {
	if l < 0 || int(l) > len(levelStrings) {
		return "UNKNOWN"
	}
	return levelStrings[int(l)]
}
