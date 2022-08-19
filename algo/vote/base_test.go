package vote_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

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
