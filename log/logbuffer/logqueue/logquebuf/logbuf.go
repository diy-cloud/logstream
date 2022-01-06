package logquebuf

import (
	"errors"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/logbuffer"
	"github.com/snowmerak/logstream/log/logbuffer/logqueue"
)

type LogQueueBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
	signals    map[string]chan struct{}
}

func New(bufferSize int) *LogQueueBuffer {
	return &LogQueueBuffer{
		trie:       ctrie.New(nil),
		bufferSize: bufferSize,
		signals:    map[string]chan struct{}{},
	}
}

func (e *LogQueueBuffer) AddTopic(topic string, signal chan struct{}) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, logqueue.New(e.bufferSize))
	}
	if _, ok := e.signals[topic]; !ok {
		e.signals[topic] = signal
	}
}

func (e *LogQueueBuffer) RemoveTopic(topic string) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); ok {
		e.trie.Remove(key)
	}
	delete(e.signals, topic)
}

func (e *LogQueueBuffer) EnQueue(topic string, value log.Log) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		e.trie.Insert(key, logqueue.New(e.bufferSize))
	}
	p, _ := e.trie.Lookup(key)
	logBuffer := p.(logbuffer.LogBuffer)
	logBuffer.Push(value)
	if e.signals[topic] != nil {
		e.signals[topic] <- struct{}{}
	}
}

func (e *LogQueueBuffer) DeQueue(topic string) (log.Log, error) {
	key := []byte(topic)
	if _, ok := e.trie.Lookup(key); !ok {
		return log.Log{}, errors.New("LogBuffer.DeQueue: topic not found")
	}
	p, _ := e.trie.Lookup(key)
	logBuffer := p.(logbuffer.LogBuffer)
	return logBuffer.Pop()
}
