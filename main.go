package msgbuf

import (
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/msgbuf/unlock"
)

type MessageBuffer[T any] struct {
	trie       *ctrie.Ctrie
	bufferSize int
}

func New[T any](bufferSize int) MessageBuffer[T] {
	return MessageBuffer[T]{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
	}
}

func (e *MessageBuffer[T]) EnQueue(topic string, value T) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewRingBuffer[T](e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.RingBuffer[T])
	ringBuffer.EnQueue(value)
}

func (e *MessageBuffer[T]) DeQueue(topic string) (value T, exists bool) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return value, false
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.RingBuffer[T])
	return ringBuffer.DeQueue(), true
}
