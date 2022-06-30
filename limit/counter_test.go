package limit

import (
	"fmt"
	"testing"
	"time"
)

func Test_Counter(t *testing.T) {
	c := Counter{}
	c.Set(20, time.Second)
	reqTime := 2 * time.Second                     // 总请求时间
	reqNum := 200                                  // 总请求次数
	reqInterval := reqTime / time.Duration(reqNum) // 每次请求间隔
	var trueCount, falseCount int
	for i := 0; i < reqNum; i++ {
		go func() {
			if c.Allow() {
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

// === RUN   Test_Counter
// true count:  44
// false count:  156
// --- PASS: Test_Counter (2.07s)
