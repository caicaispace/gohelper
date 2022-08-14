package vote

import (
	"math"
	"time"

	"github.com/caicaispace/gohelper/mathx"
)

// stack overflow
// https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_stack_overflow.html

type StackOverflow struct{}

func NewStackOverflow() *StackOverflow {
	return &StackOverflow{}
}

/**
 * Qviews（问题的浏览次数）
 * Qanswers（回答的数量）
 * Qscore（问题得分）
 * Ascores（回答得分）
 * dateAsk（问题发表的时间）
 * dateActive（最后一个回答时间）
 */
func (so *StackOverflow) Hot(Qviews, Qanswers int, Qscore, Ascores float64, dateAsk, dateActive time.Time) float64 {
	now := time.Now()
	// Qage（距离问题发表的时间）
	Qage := float64(now.Unix() - dateAsk.Unix())
	Qage2, _ := mathx.MathRound(Qage/float64(3600), 1)

	// Qupdated（距离最后一个回答的时间）
	Qupdated := float64(now.Unix() - dateActive.Unix())
	Qupdated2, _ := mathx.MathRound(Qupdated/float64(3600), 1)

	dividend := (math.Log10(float64(Qviews)) * 4) + ((float64(Qanswers) * Qscore) / 5) + float64(Ascores)
	divisor := math.Pow(((Qage2 + 1) - (Qage2-Qupdated2)/2), 1.5)

	return dividend / divisor
}
