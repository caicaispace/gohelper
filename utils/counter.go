package utils

import "sync/atomic"

type counter struct {
	count int64
}

// NewCounter 计数器
func NewCounter() *counter {
	return &counter{}
}

func (c *counter) Incr(incrNum int64) {
	if incrNum <= 0 {
		incrNum = 1
	}
	atomic.AddInt64(&c.count, incrNum)
}

func (c *counter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}
