package logringbuf

import (
	"errors"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/logqueue"
	"github.com/snowmerak/logstream/unlock"
)

type LogRingBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
	signals    map[string]chan struct{}
}

func New(bufferSize int) *LogRingBuffer {
	return &LogRingBuffer{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
		signals:    map[string]chan struct{}{},
	}
}

func (e *LogRingBuffer) AddTopic(topic string, signal chan struct{}) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, unlock.NewLogRingBuffer(e.bufferSize))
	}
	if _, ok := e.signals[topic]; !ok {
		e.signals[topic] = signal
	}
}

func (e *LogRingBuffer) RemoveTopic(topic string) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); ok {
		e.trie.Remove(key)
	}
	delete(e.signals, topic)
}

func (e *LogRingBuffer) EnQueue(topic string, value log.Log) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, logqueue.New(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.LogRingBuffer)
	ringBuffer.EnQueue(value)
	if e.signals[topic] != nil {
		e.signals[topic] <- struct{}{}
	}
}

func (e *LogRingBuffer) DeQueue(topic string) (log.Log, error) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return log.Log{}, errors.New("LogBuffer.DeQueue: topic not found")
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(*unlock.LogRingBuffer)
	return ringBuffer.DeQueue(), nil
}
