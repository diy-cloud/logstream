package stdout

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"

	"github.com/snowmerak/logstream/consumer"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
)

type Stdout struct {
	sync.Mutex
	level     int32
	writer    *bufio.Writer
	converter func(log.Log) string
	ctx       context.Context
}

func New(ctx context.Context, level int32, converter func(log.Log) string) consumer.Consumer {
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
			now := time.Unix(value.UnixTime, 0)
			s.writer.Write([]byte(now.Format(time.RFC3339)))
			s.writer.Write([]byte(" "))
			switch value.Level {
			case loglevel.Debug:
				s.writer.Write([]byte("\033[0;36m"))
				s.writer.Write([]byte("[DEBUG]"))
			case loglevel.Info:
				s.writer.Write([]byte("\033[0;32m"))
				s.writer.Write([]byte("[INFO]"))
			case loglevel.Warn:
				s.writer.Write([]byte("\033[0;33m"))
				s.writer.Write([]byte("[WARN]"))
			case loglevel.Error:
				s.writer.Write([]byte("\033[0;31m"))
				s.writer.Write([]byte("[ERROR]"))
			case loglevel.Fatal:
				s.writer.Write([]byte("\033[0;35m"))
				s.writer.Write([]byte("[FATAL]"))
			default:
				s.writer.Write([]byte("\033[0;37m"))
				s.writer.Write([]byte("[UNKNOWN]"))
			}
			s.writer.Write([]byte("\033[0m"))
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
