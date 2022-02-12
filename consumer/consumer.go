package consumer

import "github.com/snowmerak/logstream/v2/log"

type Consumer interface {
	Write(log.Log) error
	Close() error
}
