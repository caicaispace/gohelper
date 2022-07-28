package mathx

import "math"

// 取两整数较小值
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 取两整数较大值
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
