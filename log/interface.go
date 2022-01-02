package log

import "github.com/snowmerak/msgbuf/log/loglevel"

type Writable interface {
	Write(level loglevel.LogLevel, value []byte) error
	Close() error
}

type Readable interface {
	Read(start, end int64) ([]string, error)
	Close() error
}
