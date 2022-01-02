package log

import (
	"strconv"
	"strings"
	"time"

	"github.com/snowmerak/msgbuf/log/loglevel"
)

type LogFactory struct {
	Time    int64
	Message strings.Builder
	Level   loglevel.LogLevel
}

type Log struct {
	Message string
	Level   loglevel.LogLevel
	Time    int64
}

func New(level loglevel.LogLevel, message string) *LogFactory {
	l := &LogFactory{}
	l.Time = time.Now().UnixMicro()
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
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + value)
	return l
}

func (l *LogFactory) AddParamInt(key string, value int) *LogFactory {
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatInt(int64(value), 10))
	return l
}

func (l *LogFactory) AddParamUint(key string, value uint) *LogFactory {
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatUint(uint64(value), 10))
	return l
}

func (l *LogFactory) AddParamBool(key string, value bool) *LogFactory {
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatBool(value))
	return l
}

func (l *LogFactory) AddParamFloat(key string, value float64) *LogFactory {
	l.Message.WriteString(" \033[0;90m" + key + "\033[0m=" + strconv.FormatFloat(value, 'f', -1, 64))
	return l
}

func (l *LogFactory) AddParamComplex(key string, value complex128) *LogFactory {
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
