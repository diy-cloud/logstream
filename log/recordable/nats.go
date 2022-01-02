package recordable

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/snowmerak/msgbuf/log"
	"github.com/snowmerak/msgbuf/log/loglevel"
)

type Nats struct {
	conn    *nats.Conn
	subject string
	level   loglevel.LogLevel
}

func NewNatsConnection(url string, subject string) (log.Writable, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Nats{
		conn:    nc,
		subject: subject,
	}, nil
}

func (n *Nats) Write(level loglevel.LogLevel, value []byte) error {
	if loglevel.Available(n.level, level) {
		return n.conn.Publish(n.subject, value)
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
