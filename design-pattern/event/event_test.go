package event

import "testing"

func TestEventFactory(t *testing.T) {
	factory := &Factory{}
	e := factory.Create(Start)
	if e.EventType() != Start {
		t.Fatalf("expect %v, but got %v", Start, e.EventType())
	}

	e = factory.Create(End)
	if e.EventType() != End {
		t.Fatalf("expect %v, but got %v", End, e.EventType())
	}

}

func TestEventFactoryDist(t *testing.T) {
	var e Event
	e = OfStart()
	if e.EventType() != Start {
		t.Fatalf("expect %v, but got %v", Start, e.EventType())
	}

	e = OfEnd()
	if e.EventType() != End {
		t.Fatalf("expect %v, but got %v", End, e.EventType())
	}

}
