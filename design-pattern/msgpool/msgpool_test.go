package msgpool

import (
	"testing"
)

func TestMessagePool(t *testing.T) {
	msg1 := Instance().GetMsg()
	expect := 0
	if msg1.Count != expect {
		t.Fatalf("expect msg Count %d, but got %d", expect, msg1.Count)
	}

	msg1.Count = 1
	Instance().AddMsg(msg1)
	msg2 := Instance().GetMsg()

	expect = 1
	if msg2.Count != expect {
		t.Fatalf("expect msg Count %d, but got %d", expect, msg2.Count)
	}
}
