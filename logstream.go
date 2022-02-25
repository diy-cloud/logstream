package logstream

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/diy-cloud/logstream/v2/log"
	"github.com/diy-cloud/logstream/v2/log/logbuffer"
	"github.com/snowmerak/gopool"
)

var receiveSignal = make(chan string, 64)

var goroutinePool = gopool.New(int64(runtime.NumCPU() * 4096))

func init() {
	go func() {
		for topic := range receiveSignal {
			t := topic
			goroutinePool.Go(
				func() interface{} {
					key := []byte(t)
					value, ok := trie.Lookup(key)
					if !ok {
						fmt.Println(time.Now().Format(time.RFC3339), "logstream: topic buffer not registered")
						return nil
					}
					buffer := value.(logbuffer.LogBuffer)
					log, err := buffer.Pop()
					if err != nil {
						fmt.Println(time.Now().Format(time.RFC3339), "logstream: topic buffer has any error: ", err.Error())
						return nil
					}
					value, ok = consumersMap.Lookup(key)
					if !ok {
						fmt.Println(time.Now().Format(time.RFC3339), "logstream: topic consumers not registered")
						return nil
					}
					consumers := value.(*Consumers)
					consumers.Lock()
					for _, consumer := range consumers.list {
						err := consumer.Write(log)
						if err != nil {
							fmt.Println(time.Now().Format(time.RFC3339), "logstream: topic consumer has any error: ", err.Error())
						}
					}
					consumers.Unlock()
					return nil
				},
			)
		}
	}()
}

func SetGoroutineMaxSize(size int64) {
	goroutinePool.SetMax(size)
}

func Write(topic string, log log.Log) error {
	key := []byte(topic)
	if _, ok := trie.Lookup(key); !ok {
		return errors.New("topic not registered")
	}

	value, ok := trie.Lookup(key)
	if !ok {
		return errors.New("topic buffer not registered")
	}
	buffer := value.(logbuffer.LogBuffer)
	buffer.Push(log)
	if !ok {
		return errors.New("topic signal not registered")
	}
	receiveSignal <- topic
	return nil
}

func Wait() {
	goroutinePool.Wait()
}
