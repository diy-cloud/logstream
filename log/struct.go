package log

import (
	"strconv"
	"strings"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/snowmerak/logstream/log/loglevel"
)

type LogFactory struct {
	Time    time.Time
	Message strings.Builder
	Level   loglevel.LogLevel

	hasParam bool
}

type Log struct {
	Message string
	Level   loglevel.LogLevel
	Time    time.Time
}

func New(level loglevel.LogLevel, message string) *LogFactory {
	l := &LogFactory{}
	l.Time = time.Now()
	l.Level = level
	switch level {
	case loglevel.Debug:
		l.Message.WriteString(loglevel.WrapColor(level, "[DEBUG] "))
	case loglevel.Info:
		l.Message.WriteString(loglevel.WrapColor(level, "[INFO] "))
	case loglevel.Warn:
		l.Message.WriteString(loglevel.WrapColor(level, "[WARN] "))
	case loglevel.Error:
		l.Message.WriteString(loglevel.WrapColor(level, "[ERROR] "))
	case loglevel.Fatal:
		l.Message.WriteString(loglevel.WrapColor(level, "[FATAL] "))
	default:
		l.Message.WriteString(loglevel.WrapColor(level, "[UNKNOWN] "))
	}
	l.Message.WriteString(message)
	return l
}

func (l *LogFactory) AddParamString(key string, value string) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + value)
	return l
}

func (l *LogFactory) AddParamInt(key string, value int) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatInt(int64(value), 10))
	return l
}

func (l *LogFactory) AddParamUint(key string, value uint) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatUint(uint64(value), 10))
	return l
}

func (l *LogFactory) AddParamBool(key string, value bool) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatBool(value))
	return l
}

func (l *LogFactory) AddParamFloat(key string, value float64) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatFloat(value, 'f', -1, 64))
	return l
}

func (l *LogFactory) AddParamComplex(key string, value complex128) *LogFactory {
	if !l.hasParam {
		l.Message.WriteString(" ?")
		l.hasParam = true
	}
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatComplex(value, 'f', -1, 64))
	return l
}

func (l *LogFactory) End() Log {
	return Log{
		Message: l.Message.String(),
		Level:   l.Level,
		Time:    l.Time,
	}
}

func (l Log) Compare(other queue.Item) int {
	o, ok := other.(Log)
	if !ok {
		return 0
	}
	if l.Time.Before(o.Time) {
		return -1
	} else if l.Time.After(o.Time) {
		return 1
	} else {
		return 0
	}
}
