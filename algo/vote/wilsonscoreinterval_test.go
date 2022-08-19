package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Test_WilsonScoreInterval(t *testing.T) {
	wsi := vote.NewWilsonScoreInterval()
	lines := make([]opts.LineData, 0)
	fmt.Println(wsi.Hot(0, 0))
	for i := 1; i <= 10; i++ {
		val := wsi.Hot(i, 10-i)
		fmt.Println(val)
		lines = append(lines, opts.LineData{Value: val})
	}

	createLineChart(lines)
}

func Test_WilsonScoreIntervalTicker(t *testing.T) {
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
		wsi := vote.NewWilsonScoreInterval()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			val := wsi.Hot(ups, 10-ups)
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
