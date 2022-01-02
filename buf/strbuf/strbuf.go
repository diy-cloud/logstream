package strbuf

import (
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/msgbuf/unlock"
)

type StringBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
}

func New(bufferSize int) StringBuffer {
	return StringBuffer{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
	}
}

func (e *StringBuffer) EnQueue(topic string, value string) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewStringRingBuffer(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.StringRingBuffer)
	ringBuffer.EnQueue(value)
}

func (e *StringBuffer) DeQueue(topic string) (value string, exists bool) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return value, false
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.StringRingBuffer)
	return ringBuffer.DeQueue(), true
}
