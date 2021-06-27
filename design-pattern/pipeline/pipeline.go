package pipeline

import (
	"design-pattern/plugin"
	"fmt"
)

type Pipeline struct {
	status plugin.Status
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	if msg != nil {
		msg = p.filter.Process(msg)
		p.output.Send(msg)
	}
}

func (p *Pipeline) Start() {
	p.output.Start()
	p.filter.Start()
	p.input.Start()
	p.status = plugin.Started

	fmt.Println("pipeline plugin started")
}

func (p *Pipeline) Stop() {
	p.input.Stop()
	p.filter.Stop()
	p.output.Stop()
	p.status = plugin.Stopped

	fmt.Println("pipeline plugin stopped")
}

func (p *Pipeline) Status() plugin.Status {
	return p.status
}
