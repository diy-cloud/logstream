package errorstream

import (
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/error-stream/unlock"
)

type ErrorStream struct {
	trie       *ctrie.Ctrie
	bufferSize int
}

func New(bufferSize int) ErrorStream {
	return ErrorStream{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
	}
}

func (e *ErrorStream) EnQueue(topic string, err error) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewRingBuffer(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.RingBuffer)
	ringBuffer.EnQueue(err)
}

func (e *ErrorStream) DeQueue(topic string) (error, bool) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return nil, false
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.RingBuffer)
	return ringBuffer.DeQueue(), true
}
