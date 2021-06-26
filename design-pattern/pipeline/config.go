package pipeline

import "design-pattern/plugin"

type Config struct {
	Name   string
	Input  plugin.Config
	Filter plugin.Config
	Output plugin.Config
}

var DefaultConfig = Config{
	Name: "pipeline",
	Input: plugin.Config{
		Name:       "hello",
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
