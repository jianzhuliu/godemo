package plugin

type Type int

type Status uint8

const (
	InputType Type = iota
	FilterType
	OutputType
)

const (
	Started Status = iota + 1
	Stopped
)

type Config struct {
	Name       string
	PluginType Type
}

type Base struct {
	status Status
}

func (b *Base) Status() Status {
	return b.status
}

func (b *Base) Init(){
}
