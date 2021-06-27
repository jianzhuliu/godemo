package plugin

import (
	"design-pattern/msg"
	"fmt"
	"reflect"
)

type HelloInput struct {
	Base
}

func (h *HelloInput) Receive() *msg.Message {
	if h.Status() != Started {
		fmt.Println("Hello input plugin is not running, input nothing.")
		return nil
	}

	return msg.Builder().
		WithHeaderItem("Content-Type", "application/json").
		WithBodyItem("Hello World").
		Builder()
}

func (h *HelloInput) Start() {
	h.status = Started
	fmt.Println("Hello input plugin started")
}

func (h *HelloInput) Stop() {
	h.status = Stopped
	fmt.Println("Hello input plugin stopped")
}

func init() {
	RegisterInput("hello", reflect.TypeOf(HelloInput{}))
}
