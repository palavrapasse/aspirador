package pkg

const (
	TRACE Level = iota
	INFO
	WARNING
	ERROR
)

var (
	levelPrefix = []string{"TRACE: ", "INFO: ", "WARNING: ", "ERROR: "}
)

type Level int

func (l Level) String() string {
	if l < 0 || int(l) > len(levelPrefix) {
		l = 0
	}
	return levelPrefix[int(l)]
}
