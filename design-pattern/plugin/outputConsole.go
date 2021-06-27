package plugin

import (
	"design-pattern/msg"
	"fmt"
	"reflect"
)

type ConsoleOutput struct {
	Base
}

func (c *ConsoleOutput) Send(msg *msg.Message) {
	if c.Status() != Started {
		fmt.Println("console output plugin is not running, output nothing.")
		return
	}

	fmt.Printf("Output:\nHeader:%+v, Body:%+v\n", msg.Header.Items, msg.Body.Items)
}

func (c *ConsoleOutput) Start() {
	c.status = Started
	fmt.Println("console output plugin started")
}

func (c *ConsoleOutput) Stop() {
	c.status = Stopped
	fmt.Println("console output plugin stopped")
}

func init() {
	RegisterOutput("console", reflect.TypeOf(ConsoleOutput{}))
}
