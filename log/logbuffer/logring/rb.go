package logring

import (
	"fmt"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/snowmerak/logstream/v2/log"
	"github.com/snowmerak/logstream/v2/log/logbuffer"
)

type LogRingBuffer struct {
	ringbuffer *queue.RingBuffer
}

func New(size int) logbuffer.LogBuffer {
	return &LogRingBuffer{
		ringbuffer: queue.NewRingBuffer(uint64(size)),
	}
}

func (lrb *LogRingBuffer) Push(value log.Log) error {
	return lrb.ringbuffer.Put(value)
}

func (lrb *LogRingBuffer) Pop() (log.Log, error) {
	item, err := lrb.ringbuffer.Get()
	if err != nil {
		return log.Log{}, fmt.Errorf("LogRingBuffer.Pop: %w", err)
	}
	return item.(log.Log), nil
}

func (lrb *LogRingBuffer) Size() int {
	return int(lrb.ringbuffer.Len())
}
