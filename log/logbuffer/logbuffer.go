package logbuffer

import "github.com/snowmerak/logstream/v2/log"

type LogBuffer interface {
	Push(log log.Log) error
	Pop() (log.Log, error)
	Size() int
}
