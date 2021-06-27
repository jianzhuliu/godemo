package plugin

import "design-pattern/msg"

type Plugin interface {
	Start()
	Stop()
	Status() Status
	Init()
}

type Input interface {
	Plugin
	Receive() *msg.Message
}

type Filter interface {
	Plugin
	Process(msg *msg.Message) *msg.Message
}

type Output interface {
	Plugin
	Send(msg *msg.Message)
}

type Factory interface {
	Create(conf Config) Plugin
}
