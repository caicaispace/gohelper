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
// https://www.codeleading.com/article/86844186949/

// T_now：当前温度
// T_last: 上次温度
// tx：与上次测量的时间间隔
// coefficient: 冷却系数

// T_now = T_last * Exp(-(tx) * coefficient)

// 本期得分 = 上一期得分 x exp(-(冷却系数) x 间隔的小时数)

const (
	decayTime = 1 // 衰退时间
)

type NLOC struct {
	lastScore float64
}

func NewNLOC() *NLOC {
	return &NLOC{}
}

func (n *NLOC) Hot(timex time.Time) float64 {
	td := time.Now().Unix() - timex.Unix()
	if td < 0 {
		td = 0
	}
	if n.lastScore == 0 {
		n.lastScore = 1
	}
	thisScore := n.lastScore * math.Exp(float64(-td)*float64(decayTime))
	n.lastScore = thisScore
	return thisScore
}
