package recordable

import (
	"bufio"
	"os"
	"sync"
	"time"

	"github.com/snowmerak/msgbuf/log"
	"github.com/snowmerak/msgbuf/log/loglevel"
)

type Stdout struct {
	sync.Mutex
	level       loglevel.LogLevel
	writer      *bufio.Writer
	displayTime bool
}

func NewStdout(level loglevel.LogLevel, displayTime bool) log.Writable {
	return &Stdout{
		writer:      bufio.NewWriter(os.Stdout),
		level:       level,
		displayTime: displayTime,
	}
}

func (s *Stdout) Write(level loglevel.LogLevel, value []byte) error {
	s.Lock()
	defer s.Unlock()
	if loglevel.Available(s.level, level) {
		if s.displayTime {
			s.writer.Write([]byte(time.Now().Format(time.RFC3339)))
			s.writer.Write([]byte(" "))
		}
		s.writer.Write(value)
		s.writer.WriteByte('\n')
		return s.writer.Flush()
	}
	return nil
}

func (s *Stdout) Close() error {
	return nil
}
