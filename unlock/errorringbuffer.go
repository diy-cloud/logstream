package unlock

import (
	"runtime"
	"sync/atomic"
)

type ErrorRingBuffer struct {
	buf     []error
	size    int
	r, w    int
	counter int64
	TLock
}

func NewErrorRingBuffer(size int) *ErrorRingBuffer {
	r := new(ErrorRingBuffer)
	r.buf = make([]error, size)
	r.size = size
	return r
}

func (b *ErrorRingBuffer) EnQueue(x error) {
	for {
		ctr := atomic.LoadInt64(&b.counter)
		if ctr+1 > int64(b.size) {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr+1) {
			break
		}
	}
	b.Lock()
	b.buf[b.w] = x
	b.w++
	if b.w >= b.size {
		b.w = 0
	}
	b.Unlock()
}

func (b *ErrorRingBuffer) DeQueue() error {
	for {
		ctr := atomic.LoadInt64(&b.counter)
		if ctr <= 0 {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr-1) {
			break
		}
	}
	b.Lock()
	val := b.buf[b.r]
	b.r++
	if b.r >= b.size {
		b.r = 0
	}
	b.Unlock()
	return val
}

func (b *ErrorRingBuffer) EnQueueMany(x []error) {
	length := len(x)
	for {
		ctr := atomic.LoadInt64(&b.counter)
		if ctr+int64(length) > int64(b.size) {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr+int64(length)) {
			break
		}
	}
	b.Lock()
	for i := range x {
		b.buf[b.w] = x[i]
		b.w++
		if b.w >= b.size {
			b.w = 0
		}
	}
	b.Unlock()
}

func (b *ErrorRingBuffer) DeQueueMany(dst []error) {
	length := len(dst)
	for {
		ctr := atomic.LoadInt64(&b.counter)
		if ctr < int64(length) {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr-int64(length)) {
			break
		}
	}
	b.Lock()
	for i := range dst {
		dst[i] = b.buf[b.r]
		b.r++
		if b.r >= b.size {
			b.r = 0
		}
	}
	b.Unlock()
}