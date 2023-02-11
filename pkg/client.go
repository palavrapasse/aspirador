package pkg

type Client interface {
	Write(ar Record)
	SupportsLevel(l Level) bool
}
