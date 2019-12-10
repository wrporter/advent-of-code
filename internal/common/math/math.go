package math

import "math"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
