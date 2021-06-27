/*
建造者模式 (Builder Pattern)
*/
package msg

import "sync"

type Message struct {
	Header *Header
	Body   *Body
}

type Header struct {
	SrcAddr  string
	SrcPort  int
	DestAddr string
	DestPort int
	Items    map[string]string
}

type Body struct {
	Items []string
}

type builder struct {
	msg  *Message
	once *sync.Once
}

func Builder() *builder {
	return &builder{
		msg: &Message{
			Header: &Header{},
			Body:   &Body{},
		},
		once: &sync.Once{},
	}
}

func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

func (b *builder) WithSrcPort(srcPort int) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}

func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}

func (b *builder) WithDestPort(destPort int) *builder {
	b.msg.Header.DestPort = destPort
	return b
}

func (b *builder) WithHeaderItem(key, value string) *builder {
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})

	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(item string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, item)
	return b
}

func (b *builder) WithBodyItems(items []string) *builder {
	b.msg.Body.Items = items
	return b
}

func (b *builder) Builder() *Message {
	return b.msg
}
