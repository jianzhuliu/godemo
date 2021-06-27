package kafka

type Records struct {
	Items []string
}

type Consumer interface {
	Poll() *Records
}
