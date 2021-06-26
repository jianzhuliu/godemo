package prototype

import "testing"

func TestPrototype(t *testing.T) {
	msg := &Message{
		Count: 1,
	}

	msgCopy := msg.Clone().(*Message)
	if msgCopy.Count != msg.Count {
		t.Fatalf("expect %d, but got %d", msg.Count, msgCopy.Count)
	}

}
