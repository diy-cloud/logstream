package stdout

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"

	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
	"github.com/snowmerak/logstream/log/writable"
)

type Stdout struct {
	sync.Mutex
	level     loglevel.LogLevel
	writer    *bufio.Writer
	converter func(log.Log) string
	ctx       context.Context
}

func New(ctx context.Context, level loglevel.LogLevel, converter func(log.Log) string) writable.Writable {
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
