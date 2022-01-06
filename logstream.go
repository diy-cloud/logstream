package logstream

import (
	"context"
	"errors"
	"sync"

	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
	"github.com/snowmerak/logstream/log/writable"
)

type writer struct {
	list   []writable.Writable
	signal chan struct{}
}

type GlobalQueue struct {
	ctx     context.Context
	buf     LogStream
	writers map[string]writer
	lock    *sync.Mutex
	bufSize int
}

type LogStream interface {
	AddTopic(topic string, signal chan struct{})
	RemoveTopic(topic string)
	EnQueue(topic string, value log.Log)
	DeQueue(topic string) (log.Log, error)
}

func New(ctx context.Context, buf LogStream, bufSize int) *GlobalQueue {
	return &GlobalQueue{
		ctx:     ctx,
		buf:     buf,
		writers: map[string]writer{},
		lock:    &sync.Mutex{},
		bufSize: bufSize,
	}
}

func (ls *GlobalQueue) ObserveTopic(topic string, writers ...writable.Writable) error {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	if _, ok := ls.writers[topic]; ok {
		return errors.New("LogStream.AddTopic: topic already exists")
	}
	ls.writers[topic] = writer{
		list:   writers,
		signal: make(chan struct{}, ls.bufSize),
	}
	ls.buf.AddTopic(topic, ls.writers[topic].signal)
	go func() {
		for {
			select {
			case <-ls.ctx.Done():
				ls.lock.Lock()
				for _, w := range ls.writers[topic].list {
					w.Close()
				}
				ls.buf.RemoveTopic(topic)
				close(ls.writers[topic].signal)
				delete(ls.writers, topic)
				ls.lock.Unlock()
				return
			case <-ls.writers[topic].signal:
				l, err := ls.buf.DeQueue(topic)
				if err != nil {
					l = log.New(loglevel.Fatal, err.Error()).End()
				}
				for _, w := range ls.writers[topic].list {
					w.Write(l)
				}
			}
		}
	}()
	return nil
}

func (ls *GlobalQueue) Write(topic string, l log.Log) {
	ls.buf.EnQueue(topic, l)
}
