package vote_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func Test_NLOC(t *testing.T) {
	timer := time.NewTimer(time.Second * 10)
	quit := make(chan struct{})

	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()
	items := make([]opts.LineData, 0)

	timex := time.Now()
	go func() {
		n := vote.NewNLOC()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			val := n.Hot(timex)
			fmt.Println(val)
			items = append(items, opts.LineData{Value: val})
		}
	}()
	for {
		<-quit
		fmt.Println(items)
		createLineChart(items)
		return
	}
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", value), 64)
	return value
}

func createLineChart(LineData []opts.LineData) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line chart in Go",
			Subtitle: "This is fun to use!",
		}),
	)

	xAxis := make([]string, 0)
	for k := range LineData {
		xAxis = append(xAxis, strconv.Itoa(k+1))
	}

	// Put data into instance
	line.SetXAxis(xAxis).
		AddSeries("Category A", LineData).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create("line.html")
	_ = line.Render(f)
}
