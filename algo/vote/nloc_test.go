package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Test_NLOC(t *testing.T) {
	timer := time.NewTimer(time.Second * 10)
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	lines := make([]opts.LineData, 0)
	timex := time.Now()
	go func() {
		n := vote.NewNLOC()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			val := n.Hot(timex)
			fmt.Println(val)
			lines = append(lines, opts.LineData{Value: val})
		}
	}()

	for {
		<-quit
		createLineChart(lines)
		return
	}
}
