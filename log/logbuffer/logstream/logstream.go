package logstream

import (
	"errors"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/logbuffer"
	"github.com/snowmerak/logstream/log/logbuffer/logqueue"
)

type LogStream struct {
	trie              *ctrie.Ctrie
	bufferSize        int
	signals           map[string]chan struct{}
	bufferConstructor func(int) logbuffer.LogBuffer
}

func New(bufferSize int, bufferConstructor func(int) logbuffer.LogBuffer) *LogStream {
	return &LogStream{
		trie:              ctrie.New(nil),
		bufferSize:        bufferSize,
		signals:           map[string]chan struct{}{},
		bufferConstructor: bufferConstructor,
	}
}

func (e *LogStream) AddTopic(topic string, signal chan struct{}) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, e.bufferConstructor(e.bufferSize))
	}
	if _, ok := e.signals[topic]; !ok {
		e.signals[topic] = signal
	}
}

func (e *LogStream) RemoveTopic(topic string) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); ok {
		e.trie.Remove(key)
	}
	delete(e.signals, topic)
}

func (e *LogStream) EnQueue(topic string, value log.Log) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, logqueue.New(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(logbuffer.LogBuffer)
	ringBuffer.Push(value)
	if e.signals[topic] != nil {
		e.signals[topic] <- struct{}{}
	}
}

func (e *LogStream) DeQueue(topic string) (log.Log, error) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return log.Log{}, errors.New("LogBuffer.DeQueue: topic not found")
	}
	p, _ := e.trie.Lookup(key)
	ringBuffer := p.(logbuffer.LogBuffer)
	return ringBuffer.Pop()
}
