package vote_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Test_StackOverflow(t *testing.T) {
	so := vote.NewStackOverflow()
	dateAsk := time.Now()
	lines := make([]opts.LineData, 0)
	for i := 1; i <= 10; i++ {
		val := so.Hot(i, 10, 3, 3, dateAsk, time.Now())
		fmt.Println(val)
		lines = append(lines, opts.LineData{Value: val})
	}

	createLineChart(lines)
}

func Test_StackOverflowTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 10)
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	lines := make([]opts.LineData, 0)
	dateAsk := time.Now()
	Qviews := 1
	go func() {
		so := vote.NewStackOverflow()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			val := so.Hot(int(math.Exp(float64(Qviews))), 10, 3, 3, dateAsk, time.Now())
			fmt.Println(val)
			lines = append(lines, opts.LineData{Value: val})
			Qviews++
		}
	}()

	for {
		<-quit
		createLineChart(lines)
		return
	}
}
