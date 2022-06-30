package limit

import (
	"fmt"
	"testing"
	"time"
)

func Test_ChannelTokenBucket(t *testing.T) {
	ctb := NewChannelTokenBucket(1000, 10)
	reqTime := 2 * time.Second                     // 总请求时间
	reqNum := 200                                  // 总请求次数
	reqInterval := reqTime / time.Duration(reqNum) // 每次请求间隔
	var trueCount, falseCount int
	for i := 0; i < reqNum; i++ {
		go func() {
			if ctb.Allow() {
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
