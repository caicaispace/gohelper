package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Test_Reddit(t *testing.T) {
	reddit := vote.NewReddit()
	lines := make([]opts.LineData, 0)
	postTime := time.Now()

	for i := 1; i < 10; i++ {
		val := reddit.Hot(i*2, 10-(i/10), postTime)
		fmt.Println(val)
		lines = append(lines, opts.LineData{Value: val})
	}

	createLineChart(lines)
}

func Test_RedditTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 10)
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	lines := make([]opts.LineData, 0)
	ups := 1
	go func() {
		reddit := vote.NewReddit()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			val := reddit.Hot(ups, 10-ups, time.Now())
			fmt.Println(val)
			lines = append(lines, opts.LineData{Value: val})
			ups++
		}
	}()

	for {
		<-quit
		createLineChart(lines)
		return
	}
}
