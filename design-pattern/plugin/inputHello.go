package plugin

import "reflect"

type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

func init() {
	RegisterInput("hello", reflect.TypeOf(HelloInput{}))
}
