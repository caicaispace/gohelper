package vote

// https://github.com/VividCortex/ewma/blob/master/ewma.go
// https://blog.csdn.net/mzpmzk/article/details/80085929

// Package ewma implements exponentially weighted moving averages.

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

const (
	// By default, we average over a one-minute period, which means the average
	// age of the metrics in the period is 30 seconds.
	// 默认情况下，我们以一分钟为一个周期进行平均，这意味着该周期内指标的平均
	// 这意味着这段时间内指标的平均年龄是30秒。
	AVG_METRIC_AGE float64 = 30.0

	// The formula for computing the decay factor from the average age comes
	// from "Production and Operations Analysis" by Steven Nahmias.
	// 从平均年龄计算衰减系数的公式来自于
	// 来自史蒂文-纳米亚斯的《生产与运营分析》。
	DECAY float64 = 2 / (float64(AVG_METRIC_AGE) + 1)

	// For best results, the moving average should not be initialized to the
	// samples it sees immediately. The book "Production and Operations
	// Analysis" by Steven Nahmias suggests initializing the moving average to
	// the mean of the first 10 samples. Until the VariableEwma has seen this
	// many samples, it is not "ready" to be queried for the value of the
	// moving average. This adds some memory cost.
	// 为了获得最佳效果，移动平均线不应该被初始化为它所看到的
	// 它立即看到的样本。Steven Nahmias所著的《生产与经营
	// 中建议将移动平均线初始化为
	// 前10个样本的平均值。在VariableEwma看到这么多样本之前
	// 许多样本，它还没有 "准备好 "被查询到移动平均线的值。
	// 移动平均数。这增加了一些内存成本。
	WARMUP_SAMPLES uint8 = 10
)

// MovingAverage is the interface that computes a moving average over a time-
// series stream of numbers. The average may be over a window or exponentially
// decaying.
// MovingAverage是一个接口，用于计算一个时间序列上的移动平均线。
// 系列的数字流。平均数可以是在一个窗口上，也可以是指数型的
// 衰减。
type MovingAverage interface {
	Add(float64)
	Value() float64
	Set(float64)
}

// NewMovingAverage constructs a MovingAverage that computes an average with the
// desired characteristics in the moving window or exponential decay. If no
// age is given, it constructs a default exponentially weighted implementation
// that consumes minimal memory. The age is related to the decay factor alpha
// by the formula given for the DECAY constant. It signifies the average age
// of the samples as time goes to infinity.
// NewMovingAverage构造了一个MovingAverage，它计算出了一个具有在移动窗口或指数衰减中的
// 所需特性。如果没有给出
// 年龄，它将构造一个默认的指数加权的实现。
// 消耗最小的内存。年龄与衰减因子α有关。
// 通过给定的DECAY常数的公式。它标志着样本的平均年龄
// 它标志着样本的平均年龄，随着时间的推移达到无限大。
func NewMovingAverage(age ...float64) MovingAverage {
	if len(age) == 0 || age[0] == AVG_METRIC_AGE {
		return new(SimpleEWMA)
	}
	return &VariableEWMA{
		decay: 2 / (age[0] + 1),
	}
}

// A SimpleEWMA represents the exponentially weighted moving average of a
// series of numbers. It WILL have different behavior than the VariableEWMA
// for multiple reasons. It has no warm-up period and it uses a constant
// decay.  These properties let it use less memory.  It will also behave
// differently when it's equal to zero, which is assumed to mean
// uninitialized, so if a value is likely to actually become zero over time,
// then any non-zero value will cause a sharp jump instead of a small change.
// However, note that this takes a long time, and the value may just
// decays to a stable value that's close to zero, but which won't be mistaken
// for uninitialized. See http://play.golang.org/p/litxBDr_RC for example.
// 简单EWMA表示一个数字的指数加权移动平均值。
// 的一系列数字。由于多种原因，它的行为会与变量EWMA不同。
// 有多种原因。它没有预热期，并且使用一个恒定的
// 衰减。 这些特性让它使用更少的内存。 它的行为也会
// 当它等于零时，会有不同的表现，这被认为是指
// 未被初始化，所以如果一个值很可能随着时间的推移实际变成零。
// 那么任何非零值都会导致一个急剧的跳跃，而不是一个小的变化。
// 然而，请注意，这需要很长的时间，而且该值可能只是
// 衰减到一个接近零的稳定值，但不会被误认为
// 为未初始化的。请看http://play.golang.org/p/litxBDr_RC 为例。
type SimpleEWMA struct {
	// The current value of the average. After adding with Add(), this is
	// updated to reflect the average of all values seen thus far.
	// 当前平均值。在用Add()添加后，这个值会更新所有数值的平均值。
	value float64
}

// Add adds a value to the series and updates the moving average.
// 添加并更新滑动平均值
func (e *SimpleEWMA) Add(value float64) {
	if e.value == 0 { // this is a proxy for "uninitialized"
		e.value = value
	} else {
		e.value = (value * DECAY) + (e.value * (1 - DECAY))
	}
}

// Value returns the current value of the moving average.
// 获取当前滑动平均值
func (e *SimpleEWMA) Value() float64 {
	return e.value
}

// Set sets the EWMA's value.
// 设置 ewma 值
func (e *SimpleEWMA) Set(value float64) {
	e.value = value
}

// VariableEWMA represents the exponentially weighted moving average of a series of
// numbers. Unlike SimpleEWMA, it supports a custom age, and thus uses more memory.
type VariableEWMA struct {
	// The multiplier factor by which the previous samples decay.
	decay float64
	// The current value of the average.
	value float64
	// The number of samples added to this instance.
	count uint8
}

// Add adds a value to the series and updates the moving average.
func (e *VariableEWMA) Add(value float64) {
	switch {
	case e.count < WARMUP_SAMPLES:
		e.count++
		e.value += value
	case e.count == WARMUP_SAMPLES:
		e.count++
		e.value = e.value / float64(WARMUP_SAMPLES)
		e.value = (value * e.decay) + (e.value * (1 - e.decay))
	default:
		e.value = (value * e.decay) + (e.value * (1 - e.decay))
	}
}

// Value returns the current value of the average, or 0.0 if the series hasn't
// warmed up yet.
func (e *VariableEWMA) Value() float64 {
	if e.count <= WARMUP_SAMPLES {
		return 0.0
	}

	return e.value
}

// Set sets the EWMA's value.
func (e *VariableEWMA) Set(value float64) {
	e.value = value
	if e.count <= WARMUP_SAMPLES {
		e.count = WARMUP_SAMPLES + 1
	}
}
