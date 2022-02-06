package log

import (
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/snowmerak/log-silo/log"
)

type Log log.Log

func New(appID int32, level int32, msg string, param ...Param) *Log {
	l := new(Log)
	l.AppID = appID
	l.Level = level
	l.UnixTime = time.Now().Unix()

	length := len(msg) + 3
	for _, p := range param {
		length += len(p) + 1
	}
	buf := make([]byte, length)

	cur := 0
	copy(buf[cur:], msg)
	cur += len(msg)
	copy(buf[cur:], " ? ")
	cur += 3
	for _, p := range param {
		copy(buf[cur:], p)
		cur += len(p)
		copy(buf[cur:], " ")
		cur += 1
	}

	l.Message = string(buf)

	return l
}

func (l Log) Compare(o queue.Item) int {
	otherLog, ok := o.(Log)
	if !ok {
		return 0
	}
	if l.UnixTime < otherLog.UnixTime {
		return -1
	}
	if l.UnixTime > otherLog.UnixTime {
		return 1
	}
	return 0
}

type Param string

func String(k string, v string) Param {
	return Param(k + "=\"" + v + "\"")
}

func Int(k string, v int) Param {
	return Param(k + "=" + strconv.FormatInt(int64(v), 10))
}

func Int8(k string, v int8) Param {
	return Param(k + "=" + strconv.FormatInt(int64(v), 10))
}

func Int16(k string, v int16) Param {
	return Param(k + "=" + strconv.FormatInt(int64(v), 10))
}

func Int32(k string, v int32) Param {
	return Param(k + "=" + strconv.FormatInt(int64(v), 10))
}

func Int64(k string, v int64) Param {
	return Param(k + "=" + strconv.FormatInt(v, 10))
}

func Uint(k string, v uint) Param {
	return Param(k + "=" + strconv.FormatUint(uint64(v), 10))
}

func Uint8(k string, v uint8) Param {
	return Param(k + "=" + strconv.FormatUint(uint64(v), 10))
}

func Uint16(k string, v uint16) Param {
	return Param(k + "=" + strconv.FormatUint(uint64(v), 10))
}

func Uint32(k string, v uint32) Param {
	return Param(k + "=" + strconv.FormatUint(uint64(v), 10))
}

func Uint64(k string, v uint64) Param {
	return Param(k + "=" + strconv.FormatUint(v, 10))
}

func Float64(k string, v float64) Param {
	return Param(k + "=" + strconv.FormatFloat(v, 'f', -1, 64))
}

func Float32(k string, v float32) Param {
	return Param(k + "=" + strconv.FormatFloat(float64(v), 'f', -1, 32))
}

func Byte(k string, v byte) Param {
	return Param(k + "='" + string(v) + "'")
}

func Rune(k string, v rune) Param {
	return Param(k + "='" + string(v) + "'")
}

func Duration(k string, v time.Duration) Param {
	return Param(k + "=" + v.String())
}

func Hex(k string, v []byte) Param {
	return Param(k + "=" + hex.EncodeToString(v))
}

func Binary(k string, v []byte) Param {
	buf := make([]string, len(v))
	for i, b := range v {
		buf[i] = strconv.FormatUint(uint64(b), 2)
	}
	return Param(k + "=" + strings.Join(buf, ""))
}

func Bool(k string, v bool) Param {
	return Param(k + "=" + strconv.FormatBool(v))
}
