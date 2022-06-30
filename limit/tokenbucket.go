package limit

import (
	"sync"
	"time"
)

// 令牌桶
type TokenBucket struct {
	rate         int64 // 固定的token放入速率, r/s
	capacity     int64 // 桶的容量
	tokens       int64 // 桶中当前token数量
	lastTokenSec int64 // 桶上次放token的时间戳 s

	lock sync.Mutex
}

// 判断是否可通过
func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().Unix()
	// 先添加初始令牌
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	}
	// 没有令牌,则拒绝
	return false
}

// 动态设置参数
// r rate
// c capacity
func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = r
	l.lastTokenSec = time.Now().Unix()
}
