package limit

import (
	"fmt"
	"testing"
	"time"
)

func Test_TokenBucket(t *testing.T) {
	lb := &TokenBucket{}
	lb.Set(20, 20)
	reqTime := 2 * time.Second                     // 总请求时间
	reqNum := 200                                  // 总请求次数
	reqInterval := reqTime / time.Duration(reqNum) // 每次请求间隔
	var trueCount, falseCount int
	for i := 0; i < reqNum; i++ {
		go func() {
			if lb.Allow() {
				trueCount++
			} else {
				falseCount++
			}
		}()
		time.Sleep(reqInterval)
	}
	fmt.Println("true count: ", trueCount)
	fmt.Println("false count: ", falseCount)
}

// === RUN   Test_TokenBucket
// true count:  60
// false count:  140
// --- PASS: Test_TokenBucket (2.07s)
