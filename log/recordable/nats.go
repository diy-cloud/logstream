package recordable

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/snowmerak/msgbuf/log"
	"github.com/snowmerak/msgbuf/log/loglevel"
)

var Conn map[string]*nats.Conn

type Nats struct {
	conn      *nats.Conn
	subject   string
	level     loglevel.LogLevel
	converter func(log.Log) string
}

func init() {
	Conn = make(map[string]*nats.Conn)
}

func NewNatsConnection(url string, subject string, converter func(log.Log) string) (log.Writable, error) {
	if _, ok := Conn[url]; !ok {
		var err error
		Conn[url], err = nats.Connect(url)
		if err != nil {
			return nil, err
		}
	}
	return &Nats{
		conn:      Conn[url],
		subject:   subject,
		converter: converter,
	}, nil
}

func (n *Nats) Write(level loglevel.LogLevel, value log.Log) error {
	if loglevel.Available(n.level, level) {
		if n.converter == nil {
			return n.conn.Publish(n.subject, []byte(value.Message))
		}
		return n.conn.Publish(n.subject, []byte(n.converter(value)))
	}
	return nil
}

func (n *Nats) Close() error {
	if err := n.conn.Flush(); err != nil {
		return fmt.Errorf("Nats.Close: %w", err)
	}
	if err := n.conn.Drain(); err != nil {
		return fmt.Errorf("Nats.Close: %w", err)
	}
	n.conn.Close()
	return nil
}
