package plugin

import (
	"design-pattern/msg"
	"fmt"
	"reflect"
	"strings"
)

type UpperFilter struct {
	Base
}

func (u *UpperFilter) Process(msg *msg.Message) *msg.Message {
	if u.Status() != Started {
		fmt.Println("upper filter plugin is not running, filter nothing.")
		return nil
	}

	for i, item := range msg.Body.Items {
		msg.Body.Items[i] = strings.ToUpper(item)
	}

	return msg
}

func (u *UpperFilter) Start() {
	u.status = Started
	fmt.Println("upper filter plugin started")
}

func (u *UpperFilter) Stop() {
	u.status = Stopped
	fmt.Println("upper filter plugin stopped")
}

func init() {
	RegisterFilter("upper", reflect.TypeOf(UpperFilter{}))
}
