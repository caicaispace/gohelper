package vote

// bayesian average (贝叶斯平均)
// https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_bayesian_average.html
type BayesianAverage struct{}

func NewBayesianAverage() *BayesianAverage {
	return &BayesianAverage{}
}
