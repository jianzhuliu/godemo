/*
原型模式 (Prototype Pattern)
*/

package prototype

type Prototype interface {
	Clone() Prototype
}

type Message struct {
	Count int
}

func (m *Message) Clone() Prototype {
	copyMsg := *m
	return &copyMsg
}
