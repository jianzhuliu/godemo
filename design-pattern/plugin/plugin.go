package plugin

type Type int

const (
	InputType Type = iota
	FilterType
	OutputType
)

type Config struct {
	Name       string
	PluginType Type
}
