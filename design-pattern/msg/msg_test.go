package msg

import (
	"reflect"
	"testing"
)

var testMessageData = []Message{
	Message{
		&Header{
			SrcAddr:  "127.0.0.1",
			SrcPort:  8080,
			DestAddr: "192.168.126.71",
			DestPort: 9090,
			Items: map[string]string{
				"Content-Type": "application/json",
			},
		},
		&Body{
			Items: []string{"<div>Hello World</div>"},
		},
	},
}

func TestMessageBuilder(t *testing.T) {
	for _, data := range testMessageData {
		msgBuilder := Builder().
			WithSrcAddr(data.Header.SrcAddr).
			WithSrcPort(data.Header.SrcPort).
			WithDestAddr(data.Header.DestAddr).
			WithDestPort(data.Header.DestPort)

		for k, v := range data.Header.Items {
			msgBuilder.WithHeaderItem(k, v)
		}

		for _, item := range data.Body.Items {
			msgBuilder.WithBodyItem(item)
		}

		msg := msgBuilder.Builder()

		if msg.Header.SrcAddr != data.Header.SrcAddr {
			t.Fatalf("Header.SrcAddr| expect %s, but got %s", data.Header.SrcAddr, msg.Header.SrcAddr)
		}

		if msg.Header.SrcPort != data.Header.SrcPort {
			t.Fatalf("Header.SrcPort| expect %d, but got %d", data.Header.SrcPort, msg.Header.SrcPort)
		}

		if msg.Header.DestAddr != data.Header.DestAddr {
			t.Fatalf("Header.DestAddr| expect %s, but got %s", data.Header.DestAddr, msg.Header.DestAddr)
		}

		if msg.Header.DestPort != data.Header.DestPort {
			t.Fatalf("Header.DestPort| expect %d, but got %d", data.Header.DestPort, msg.Header.DestPort)
		}

		if !reflect.DeepEqual(msg.Header.Items, data.Header.Items) {
			t.Fatalf("Header.Items===============\nexpect============\n%#v,\nbut got==========\n%#v", data.Header.Items, msg.Header.Items)
		}

		if !reflect.DeepEqual(msg.Body.Items, data.Body.Items) {
			t.Fatalf("Body.Items===============\nexpect============\n%#v,\nbut got==========\n%#v", data.Body.Items, msg.Body.Items)
		}

		/*
			if !reflect.DeepEqual(msg.Header, data.Header) {
				t.Fatalf("Header|expect %#v, but got %#v", data.Header, msg.Header)
			}

			if !reflect.DeepEqual(msg.Body, data.Body) {
				t.Fatalf("Body|expect %#v, but got %#v", data.Body, msg.Body)
			}
			//*/
	}
}
