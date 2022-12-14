package datetime

import (
	"time"

	"github.com/caicaispace/gohelper/utils"
)

type Timespan time.Duration

// 时间戳格式化
func (t Timespan) Format(format string) string {
	timestamp := int64(t)
	if utils.NumberLen(timestamp) >= 13 {
		timestamp = timestamp / 1000
	}
	return time.Unix(int64(timestamp), 0).Format(format)
}

// 日期转时间戳
func DatetimeToUnix(datetimeStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix() // 1504082441
}

func NowTimestampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
