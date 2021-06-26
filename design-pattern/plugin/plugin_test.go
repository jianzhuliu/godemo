package plugin

import (
	"strings"
	"testing"
)

func TestFactoryInput(t *testing.T) {
	conf := Config{
		Name:       "hello",
		PluginType: InputType,
	}

	f := &InputFactory{}
	input := f.Create(conf).(Input)

	expect := "Hello World"

	if input.Receive() != expect {
		t.Fatalf("expect %s, but got %s", expect, input.Receive())
	}
}

func TestFactoryFilter(t *testing.T) {
	conf := Config{
		Name:       "upper",
		PluginType: FilterType,
	}

	f := &FilterFactory{}
	filter := f.Create(conf).(Filter)

	msg := "Hello World"
	expect := strings.ToUpper(msg)

	if filter.Process(msg) != expect {
		t.Fatalf("expect %s, but got %s", expect, filter.Process(msg))
	}
}

func TestFactoryOutput(t *testing.T) {
	conf := Config{
		Name:       "console",
		PluginType: OutputType,
	}

	f := &OutputFactory{}
	output := f.Create(conf).(Output)

	output.Send("console from output factory")

}
