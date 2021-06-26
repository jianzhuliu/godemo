package pipeline

import "design-pattern/plugin"

type Pipeline struct {
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}
