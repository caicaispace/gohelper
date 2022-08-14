package mathx

import (
	"math"
	"strconv"
)

// MinInt 取两整数较小值
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// MaxInt 取两整数较大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const epsilon = 1e-6

// CalcEntropy calculates the entropy of m.
func CalcEntropy(m map[interface{}]int) float64 {
	if len(m) == 0 || len(m) == 1 {
		return 1
	}

	var entropy float64
	var total int
	for _, v := range m {
		total += v
	}

	for _, v := range m {
		proba := float64(v) / float64(total)
		if proba < epsilon {
			proba = epsilon
		}
		entropy -= proba * math.Log2(proba)
	}

	return entropy / math.Log2(float64(len(m)))
}

// MathRound
func MathRound(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)
}
