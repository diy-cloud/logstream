package nats

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	conn    *nats.Conn
	subject string
}

func New(url string, subject string) (*Nats, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Nats{
		conn:    nc,
		subject: subject,
	}, nil
}

func (n *Nats) Get(key string) (string, error) {
	return "", errors.New("Nats.Get: not implemented")
}

func (n *Nats) Set(key, value string) error {
	return n.conn.Publish(n.subject, []byte(value))
}

func (n *Nats) SetBytes(key []byte, value []byte) error {
	return n.conn.Publish(n.subject, value)
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
