package logbuffer

import "github.com/diy-cloud/logstream/v2/log"

type LogBuffer interface {
	Push(log log.Log) error
	Pop() (log.Log, error)
	Size() int
}
