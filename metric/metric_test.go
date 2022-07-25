package metric_test

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/logx"
	"github.com/caicaispace/gohelper/metric"
	"github.com/caicaispace/gohelper/setting"
	"github.com/caicaispace/gohelper/task"
)

var (
	metricJob          = flag.String("metric-job", "test", "prometheus job name")
	metricInstance     = flag.String("metric-instance", "", "prometheus instance name")
	metricAddress      = flag.String("metric-address", "127.0.0.1:9091", "prometheus proxy address")
	metricIntervalSync = flag.Uint64("interval-metric-sync", 1, "Interval(sec): metric sync")
	// log
	logPath      = flag.String("log-path", "logs/", "App: log file path")
	logPrefix    = flag.String("log-pref", "log_", "App: log file prefix")
	logExtension = flag.String("log-ext", "log", "App: log file extension")
)

func initAppSetting() {
	setting.App.LogPath = *logPath
	setting.App.LogPrefix = *logPrefix
	setting.App.LogExtension = *logExtension
}

func Test_New(t *testing.T) {
	initAppSetting()
	logx.Setup()

	tickerCount := 1
	conf := metric.NewMetricCfg(*metricJob, *metricInstance, *metricAddress, time.Second*time.Duration(*metricIntervalSync))
	runner := task.NewRunner()
	metric.StartMetricsPush(runner, conf)
	m := metric.NewMetric()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		tickerCount++
		if tickerCount > 10 {
			ticker.Stop()
			break
		}
		m.PostRequest("test", true, time.Time{})
		fmt.Println("get ticker", time.Now().Format("2006-01-02 15:04:05"))
	}
}
