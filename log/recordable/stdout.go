package recordable

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"

	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
)

type Stdout struct {
	sync.Mutex
	level     loglevel.LogLevel
	writer    *bufio.Writer
	converter func(log.Log) string
	waiting   []log.Log
	ctx       context.Context
}

func NewStdout(ctx context.Context, level loglevel.LogLevel, converter func(log.Log) string) log.Writable {
	s := &Stdout{
		writer:    bufio.NewWriter(os.Stdout),
		level:     level,
		converter: converter,
		ctx:       ctx,
	}
	return s
}

func (s *Stdout) Write(value log.Log) error {
	s.Lock()
	defer s.Unlock()
	if loglevel.Available(s.level, value.Level) {
		if s.converter == nil {
			s.writer.Write([]byte(value.Time.Format(time.RFC3339Nano)))
			s.writer.Write([]byte(" "))
			s.writer.Write([]byte(value.Message))
		} else {
			s.writer.Write([]byte(s.converter(value)))
		}
		s.writer.WriteByte('\n')
		s.writer.Flush()
	}
	return nil
}

func (s *Stdout) Close() error {
	return nil
}
