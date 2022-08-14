package vote

import (
	"math"
	"time"
)

// Newton's law of cooling (牛顿冷却)
// https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_newton_s_law_of_cooling.html

// 牛顿冷却定律
// https://www.khanacademy.org/math/differential-equations/first-order-differential-equations/exponential-models-diff-eq/v/applying-newtons-law-of-cooling-to-warm-oatmeal
// https://d-arora.github.io/Doing-Physics-With-Matlab/mpDocs/tp_Newton.htm

const (
	decayTime = int64(time.Millisecond * 100) // 衰退时间
)

type NLOC struct{}

func NewNLOC() *NLOC {
	return &NLOC{}
}

func (n *NLOC) Hot(timex time.Time) float64 {
	td := time.Now().Unix() - timex.Unix()
	if td < 0 {
		td = 0
	}
	w := math.Exp(float64(-td) / float64(decayTime))
	// w, _ = utils.MathRound(w, 9)
	return w
}
