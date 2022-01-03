package errbuf

import (
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/unlock"
)

type ErrorBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
}

func New(bufferSize int) ErrorBuffer {
	return ErrorBuffer{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
	}
}

func (e *ErrorBuffer) EnQueue(topic string, value error) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewErrorRingBuffer(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.ErrorRingBuffer)
	ringBuffer.EnQueue(value)
}

func (e *ErrorBuffer) DeQueue(topic string) (value error, exists bool) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return value, false
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.ErrorRingBuffer)
	return ringBuffer.DeQueue(), true
}
