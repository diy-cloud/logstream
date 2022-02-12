package logstream

import (
	"errors"
	"sync"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/snowmerak/logstream/v2/consumer"
	"github.com/snowmerak/logstream/v2/log/logbuffer"
	"github.com/snowmerak/logstream/v2/log/logbuffer/logring"
)

type Consumers struct {
	list []consumer.Consumer
	sync.Mutex
}

var bufferSize = 8
var bufferConstructor func(int) logbuffer.LogBuffer = logring.New

var trie = ctrie.New(nil)
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
