/*
单例模式 (Singleton Pattern)
1、外部不可直接实例化对象，
2、提供统一的创建实例入口

*/
package msgpool

import "sync"

type Message struct {
	Count int
}

//消息池
type messagePool struct {
	pool *sync.Pool
}

//添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

//获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}

//消息池单例
/*
//饿汉模式
var msgPool = &messagePool{
	pool: &sync.Pool{
		New: func() interface{} {
			return &Message{
				Count: 0,
			}
		},
	},
}

//提供统一访问入口
func Instance() *messagePool {
	return msgPool
}
//*/

var msgPool *messagePool
var once = &sync.Once{}

//懒汉模式，按需创建对象
func Instance() *messagePool {
	once.Do(func() {
		msgPool = &messagePool{
			pool: &sync.Pool{
				New: func() interface{} {
					return &Message{
						Count: 0,
					}
				},
			},
		}

	})

	return msgPool
}
