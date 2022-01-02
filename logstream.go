package logstream

import (
	"context"
	"fmt"
	"sync"

	"github.com/snowmerak/logstream/buf/logbuf"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
)

type LogStream struct {
	ctx     context.Context
	buf     *logbuf.LogBuffer
	writers map[string][]log.Writable
	lock    *sync.Mutex
}

func New(ctx context.Context, buf *logbuf.LogBuffer) *LogStream {
	return &LogStream{
		ctx: ctx,
		buf: buf,
	}
}

func (l *LogStream) Observe(topic string, writers ...log.Writable) error {
	if l.writers == nil {
		l.writers = make(map[string][]log.Writable)
	}
	l.writers[topic] = writers

	l.buf.AddTopic(topic)

	go func() {
		for {
			value, ok := l.buf.DeQueue(topic)
			if !ok {
				return
			}
			for _, ch := range l.writers[topic] {
				go func(ch log.Writable) {
					if err := ch.Write(value); err != nil {
						fmt.Println("LogStream.Observe: cannot write to writable on " + topic + ": " + err.Error())
					}
				}(ch)
			}
		}
	}()

	return nil
}

func (l *LogStream) Push(topic string, value log.Log) {
	l.buf.EnQueue(topic, value)
}

func (l *LogStream) CloseTopic(topic string) {
	l.buf.RemoveTopic(topic)
	l.buf.EnQueue(topic, log.Log{Level: loglevel.Fatal, Message: "LogStream.CloseTopic"})
	for _, w := range l.writers[topic] {
		if err := w.Close(); err != nil {
			fmt.Println("LogStream.CloseTopic: cannot close writable: " + err.Error())
		}
	}
	l.lock.Lock()
	l.writers[topic] = nil
	delete(l.writers, topic)
	l.lock.Unlock()
}

func (l *LogStream) Close() {
	for topic, wl := range l.writers {
		l.buf.RemoveTopic(topic)
		l.buf.EnQueue(topic, log.Log{Level: loglevel.Fatal, Message: "LogStream.Close"})
		for _, w := range wl {
			if err := w.Close(); err != nil {
				fmt.Println("LogStream.Close: cannot close writable: " + err.Error())
			}
		}
	}
	l.writers = nil
	l.buf = nil
}
