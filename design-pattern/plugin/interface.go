package plugin

type Plugin interface{}

type Input interface {
	Plugin
	Receive() string
}

type Filter interface {
	Plugin
	Process(msg string) string
}

type Output interface {
	Plugin
	Send(msg string)
}

type Factory interface {
	Create(conf Config) Plugin
}
