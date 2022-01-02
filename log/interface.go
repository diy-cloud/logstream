package log

import "github.com/snowmerak/msgbuf/log/loglevel"

type Writable interface {
	Write(level loglevel.LogLevel, value Log) error
	Close() error
}

type Readable interface {
	Read(start, end int64) ([]string, error)
	Close() error
}
