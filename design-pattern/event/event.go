/*
工厂模式 (Factory Method Pattern)
*/
package event

type EventType uint8

const (
	Start EventType = iota
	End
)

type Event interface {
	EventType() EventType
	Content() string
}

type StartEvent struct {
	content string
}

func (e *StartEvent) EventType() EventType {
	return Start
}

func (e *StartEvent) Content() string {
	return e.content
}

type EndEvent struct {
	content string
}

func (e *EndEvent) EventType() EventType {
	return End
}

func (e *EndEvent) Content() string {
	return e.content
}

type Factory struct{}

func (e *Factory) Create(etype EventType) Event {
	switch etype {
	case Start:
		return &StartEvent{content: "this is start event"}
	case End:
		return &EndEvent{content: "this is end event"}
	default:
		return nil
	}
}

func OfStart() Event {
	return &StartEvent{content: "this is start event"}
}

func OfEnd() Event {
	return &EndEvent{content: "this is end event"}
}
