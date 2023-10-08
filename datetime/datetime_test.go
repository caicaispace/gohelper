package datetime

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNowDateTime(t *testing.T) {
	fmt.Println(NowDatetime())
}

func TestFormat(t *testing.T) {
	fmt.Println(Timespan(time.Now().Unix() + 3600).Format("2006-01-02 15:04:05"))
	tt, _ := time.ParseDuration("1h")
	fmt.Println(Timespan(tt).Format("2006-01-02 15:04:05"))
	fmt.Println(Timespan(166993969).Format("2006-01-02 15:04:05"))
	fmt.Println(Timespan(1669939692000).Format("2006-01-02 15:04:05"))
	fmt.Println(strconv.Atoi(Timespan(time.Now().Unix()).Format("20060102")))
}

func TestDatetimeToUnix(t *testing.T) {
	endTime := "2022-11-14 13:30:00"
	// fmt.Println(DatetimeToUnix(endTime))
	// fmt.Println(DatetimeToUnix(endTime) + 3600)
	fmt.Println(Timespan(DatetimeToUnix(endTime) + 600).Format("2006-01-02 15:04:05"))
}

// 计算两个时间相差多少年
func TestTime2(t *testing.T) {
	fmt.Println(DateDiffYears("20200908", "20210908", "20060102"))
}
