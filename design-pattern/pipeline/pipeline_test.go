package pipeline

import "testing"

func TestPipeline(t *testing.T) {
	p := Of(DefaultConfig)
	p.Exec()
}
