package limit

import "time"

// 管道令牌桶限流
type ChannelTokenBucket struct {
	time     int // 桶清空周期
	capacity int // 桶的容量

	bucket      chan struct{}
	emptyBucket chan struct{}

	tickerCh *time.Ticker
}

func NewChannelTokenBucket(t, c int) *ChannelTokenBucket {
	cb := &ChannelTokenBucket{
		time:        t,
		capacity:    c,
		bucket:      make(chan struct{}, c),
		emptyBucket: make(chan struct{}, c),
		tickerCh:    time.NewTicker(time.Duration(t) * time.Millisecond),
	}
	go func() {
		for range cb.tickerCh.C {
			cb.clearBucket()
			// fmt.Println("Tick at", t)
		}
	}()
	return cb
}

func (cb *ChannelTokenBucket) Allow() bool {
	if len(cb.bucket) < cb.capacity {
		cb.bucket <- struct{}{}
		return true
	}
	return false
}

func (cb *ChannelTokenBucket) Set(t, c int) {
	cb.time = t
	cb.capacity = c
	cb.restTicker()
	cb.clearBucket()
}

func (cb *ChannelTokenBucket) restTicker() {
	cb.tickerCh.Stop()
	cb.tickerCh = time.NewTicker(time.Duration(cb.time) * time.Millisecond)
}

func (cb *ChannelTokenBucket) clearBucket() {
	// cb.bucket = cb.emptyBucket
	cb.bucket = make(chan struct{}, cb.capacity)
}
