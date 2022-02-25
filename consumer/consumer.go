package consumer

import "github.com/diy-cloud/logstream/v2/log"

type Consumer interface {
	Write(log.Log) error
	Close() error
}
