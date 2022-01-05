package logqueue

import (
	"errors"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/logbuffer"
)

type LogQueue struct {
	queue *queue.PriorityQueue
}

func New(size int) logbuffer.LogBuffer {
	return &LogQueue{
		queue: queue.NewPriorityQueue(size, false),
	}
}

func (lq *LogQueue) Push(log log.Log) error {
	if lq == nil {
		return errors.New("LogQueue.Push: LogQueue is nil")
	}
	return lq.queue.Put(log)
}

func (lq *LogQueue) Pop() (log.Log, error) {
	if lq == nil {
		return log.Log{}, errors.New("LogQueue.Pop: LogQueue is nil")
	}
	item, err := lq.queue.Get(1)
	if err != nil {
		return log.Log{}, err
	}
	return item[0].(log.Log), nil
}

func (lq *LogQueue) Size() int {
	if lq == nil {
		return 0
	}
	return lq.queue.Len()
}
