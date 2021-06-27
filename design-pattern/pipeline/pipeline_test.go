package pipeline

import (
	"design-pattern/plugin"
	"testing"
)

func TestPipeline(t *testing.T) {
	p := Of(DefaultConfig)
	p.Start()
	p.Exec()
	p.Stop()
}

func TestKafkaInputPipeline(t *testing.T) {
	config := Config{
		Name: "kafka input pipeline",
		Input: plugin.Config{
			Name:       "kafka",
			PluginType: plugin.InputType,
		},
		Filter: plugin.Config{
			Name:       "upper",
			PluginType: plugin.FilterType,
		},
		Output: plugin.Config{
			Name:       "console",
			PluginType: plugin.OutputType,
		},
	}

	p := Of(config)
	p.Start()
	p.Exec()
	p.Stop()
}
