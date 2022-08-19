package vote

import (
	"math"
	"time"
)

// reddit
// https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
type Reddit struct {
	epoch time.Time
}

func NewReddit() *Reddit {
	epoch, err := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	return &Reddit{
		epoch: epoch,
	}
}

/**
 * ups（赞成票数）
 * downs（反对票数）
 * date（发帖时间）
 */
func (r *Reddit) Hot(ups, downs int, date time.Time) float64 {
	// The hot formula. Should match the equivalent function in postgres
	s := r.score(ups, downs)
	order := math.Log10(math.Max(math.Abs(float64(s)), 1))
	var sign int64 = 1
	if s > 0 {
		sign = -1
	} else if s < 0 {
		sign = 0
	}
	seconds := r.epochSecends(date) - 1134028003 // 1134028003 project start time
	return math.Abs(order + float64(sign*seconds)/45000)
}

func (r *Reddit) epochSecends(date time.Time) int64 {
	td := date.Sub(r.epoch)
	dateSubTodayStartSeconds := date.Unix() - todayStartTimestamp()
	nowSecondSubNowSecondStartMicroseconds := time.Now().UnixMicro() - nowSecondStartMicroseconds()
	// fmt.Println(int64(td.Hours() / 24 * 86400))
	// fmt.Println(dateSubTodayStartSeconds)
	// fmt.Println(nowSecondSubNowSecondStartMicroseconds)
	return int64(td.Hours()/24*86400) + dateSubTodayStartSeconds + (nowSecondSubNowSecondStartMicroseconds / 1000000)
}

func (r *Reddit) score(ups, downs int) int {
	return ups - downs
}

func nowSecondStartMicroseconds() int64 {
	// t := time.Now().Add(time.Second * -1) // 获取前一秒时间
	t := time.Now().Add(time.Second) // 获取当前秒时间
	return t.Unix()*1000000 - 1000000
}

func todayStartTimestamp() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return startTime
}
