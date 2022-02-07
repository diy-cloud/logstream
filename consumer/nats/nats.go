package nats

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/nats-io/nats.go"
	"github.com/snowmerak/logstream/consumer"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
)

var connsLock = sync.Mutex{}
var conns map[string]*nats.Conn
var count map[string]*int64

type Nats struct {
	conn      *nats.Conn
	subject   string
	level     int32
	converter func(log.Log) []byte
	url       string
}

func init() {
	conns = make(map[string]*nats.Conn)
	count = make(map[string]*int64)
}

func NewNatsConnection(url string, subject string, converter func(log.Log) []byte) (consumer.Consumer, error) {
	connsLock.Lock()
	defer connsLock.Unlock()
	if _, ok := conns[url]; !ok {
		var err error
		conns[url], err = nats.Connect(url)
		if err != nil {
			return nil, err
		}
		atomic.AddInt64(count[url], 1)
	}
	return &Nats{
		conn:      conns[url],
		subject:   subject,
		converter: converter,
		url:       url,
	}, nil
}

func (n *Nats) Write(value log.Log) error {
	if loglevel.Available(n.level, value.Level) {
		if n.converter == nil {
			data, err := json.Marshal(value)
			if err != nil {
				return err
			}
			return n.conn.Publish(n.subject, data)
		}
		return n.conn.Publish(n.subject, n.converter(value))
	}
	return nil
}

func (n *Nats) Close() error {
	connsLock.Lock()
	defer connsLock.Unlock()
	atomic.AddInt64(count[n.url], -1)
	if atomic.LoadInt64(count[n.url]) > 0 {
		return nil
	}
	if err := n.conn.Flush(); err != nil {
		return fmt.Errorf("Nats.Close: %w", err)
	}
	if err := n.conn.Drain(); err != nil {
		return fmt.Errorf("Nats.Close: %w", err)
	}
	n.conn.Close()
	delete(conns, n.url)
	return nil
}
