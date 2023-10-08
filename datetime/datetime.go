package datetime

import (
	"strconv"
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

var CstSh, _ = time.LoadLocation("Asia/Shanghai") // 上海

// 获取当前时间
func NowDatetime() string {
	return time.Now().In(CstSh).Format("2006-01-02 15:04:05")
}

// 获取当前日期
func NowDate() string {
	return time.Now().In(CstSh).Format("2006-01-02")
}

// 获取当前日期(整型)
func NowDateInt() uint64 {
	date, _ := strconv.Atoi(time.Now().In(CstSh).Format("20060102"))
	return uint64(date)
}

// 获取当前时间
func NowTime() string {
	return time.Now().In(CstSh).Format("15:04:05")
}

// 获取当前时间戳(秒)
func NowTimestamp() uint64 {
	return uint64(time.Now().In(CstSh).Unix())
}

// 获取当前时间戳(毫秒)
func NowTimestampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前年份
func NowYear() int16 {
	return int16(time.Now().In(CstSh).Year())
}

// 获取当前月份
func NowMonth() uint8 {
	return uint8(time.Now().In(CstSh).Month())
}

// 获取当前日
func NowDay() uint8 {
	return uint8(time.Now().In(CstSh).Day())
}

// 获取当前小时
func NowHour() uint8 {
	return uint8(time.Now().In(CstSh).Hour())
}

// 获取当前分钟
func NowWeek() uint8 {
	return uint8(time.Now().In(CstSh).Weekday())
}

// 获取当前是一年中的第几周
func NowWeekOfYear() uint8 {
	_, week := time.Now().In(CstSh).ISOWeek()
	return uint8(week)
}

// 计算两个时间相差多少年
func DateDiffYears(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Hours() / 24 / 365)
}

// 计算两个时间相差多少月
func DateDiffMonths(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Hours() / 24 / 30)
}

// 计算两个时间相差多少天
func DateDiffDays(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Hours() / 24)
}

// 计算两个时间相差多少小时
func DateDiffHours(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Hours())
}

// 计算两个时间相差多少分钟
func DateDiffMinutes(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Minutes())
}

// 计算两个时间相差多少秒
func DateDiffSeconds(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Seconds())
}

// 计算两个时间相差多少毫秒
func DateDiffMilliseconds(start, end, timeFormat string) int {
	beginTime, _ := time.Parse(timeFormat, start) // string -> time
	endTime, _ := time.Parse(timeFormat, end)     // string -> time
	age := endTime.Sub(beginTime)
	return int(age.Milliseconds())
}
