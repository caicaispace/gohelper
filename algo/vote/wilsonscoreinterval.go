package vote

import "math"

// wilson score interval(威尔逊区间)
// https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_wilson_score_interval.html
type WilsonScoreInterval struct{}

func NewWilsonScoreInterval() *WilsonScoreInterval {
	return &WilsonScoreInterval{}
}

func (wsi *WilsonScoreInterval) Hot(ups, downs int) float64 {
	n := float64(ups + downs)

	if n == 0 {
		return 0
	}

	z := float64(1.0) // 1.44 = 85%, 1.96 = 95%
	phat := float64(ups) / n
	return ((phat + z*z/(2*n) - z*math.Sqrt((phat*(1-phat)+z*z/(4*n))/n)) / (1 + z*z/n))
}
