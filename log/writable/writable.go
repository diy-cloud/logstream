package writable

import "github.com/snowmerak/logstream/log"

type Writable interface {
	Write(log log.Log) error
	Close() error
}

type Readable interface {
	Read(start, end int64) ([]string, error)
	Close() error
}
