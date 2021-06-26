package plugin

import (
	"fmt"
	"reflect"
)

type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}

func init() {
	RegisterOutput("console", reflect.TypeOf(ConsoleOutput{}))
}
