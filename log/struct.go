package log

import (
	"strings"

	"github.com/snowmerak/msgbuf/log/loglevel"
)

type LogFactory struct {
	Message strings.Builder
	Level   loglevel.LogLevel
}

type Log struct {
	Message string
	Level   loglevel.LogLevel
}

func New(level loglevel.LogLevel, message string) *LogFactory {
	l := &LogFactory{}
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

func (l *LogFactory) AddParam(key, value string) *LogFactory {
	l.Message.WriteString(" " + key + "=" + value)
	return l
}

func (l *LogFactory) End() Log {
	return Log{
		Message: l.Message.String(),
		Level:   l.Level,
	}
}
