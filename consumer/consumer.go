package consumer

import "github.com/snowmerak/logstream/log"

type Consumer interface {
	Write(log.Log) error
	Close() error
}
