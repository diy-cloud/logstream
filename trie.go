package logstream

import (
	"errors"
	"sync"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/consumer"
	"github.com/snowmerak/logstream/log/logbuffer"
)

type Consumers struct {
	list []consumer.Consumer
	sync.Mutex
}

var bufferSize = 8
var bufferConstructor func(int) logbuffer.LogBuffer

var trie = ctrie.New(nil)
var signalMap = ctrie.New(nil)
var consumersMap = ctrie.New(nil)

type tempTrie struct{}

func (t tempTrie) SetBufferConstructor(f func(int) logbuffer.LogBuffer) {
	bufferConstructor = f
}

func (t tempTrie) SetBufferSize(size int) {
	if size < 1 {
		size = 1
	}
	bufferSize = size
}

func (t tempTrie) RegisterTopic(topic string) error {
	key := []byte(topic)
	if _, ok := trie.Lookup(key); ok {
		return errors.New("topic already registered")
	}

	trie.Insert(key, bufferConstructor(bufferSize))
	signalMap.Insert(key, make(chan struct{}, 1))
	consumers := Consumers{
		list: make([]consumer.Consumer, 0),
	}
	consumersMap.Insert(key, &consumers)

	return nil
}

func (t tempTrie) UnregisterTopic(topic string) error {
	key := []byte(topic)
	if _, ok := trie.Lookup(key); !ok {
		return errors.New("topic not registered")
	}

	trie.Remove(key)
	signalMap.Remove(key)
	consumersMap.Remove(key)

	return nil
}

func (t tempTrie) RegisterConsumer(topic string, csm consumer.Consumer) error {
	key := []byte(topic)
	if _, ok := trie.Lookup(key); !ok {
		return errors.New("topic not registered")
	}

	consumers, ok := consumersMap.Lookup(key)
	if !ok {
		cs := Consumers{
			list: make([]consumer.Consumer, 0),
		}
		consumersMap.Insert(key, &cs)
		consumers, _ = consumersMap.Lookup(key)
	}

	cs, ok := consumers.(*Consumers)
	if !ok {
		return errors.New("consumers is not a Consumers")
	}

	cs.Lock()
	cs.list = append(cs.list, csm)
	cs.Unlock()

	return nil
}

var Trie tempTrie
