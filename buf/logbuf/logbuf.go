package logbuf

import (
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/msgbuf/log"
	"github.com/snowmerak/msgbuf/unlock"
)

type LogBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
}

func New(bufferSize int) LogBuffer {
	return LogBuffer{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
	}
}

func (e *LogBuffer) EnQueue(topic string, value log.Log) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewLogRingBuffer(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.LogRingBuffer)
	ringBuffer.EnQueue(value)
}

func (e *LogBuffer) DeQueue(topic string) (value log.Log, exists bool) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return value, false
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.LogRingBuffer)
	return ringBuffer.DeQueue(), true
}
