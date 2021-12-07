package lib

import "math"

func AbsInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func MaxInInSlice(in []int) int {
	out := math.MinInt
	for _, val := range in {
		if val > out {
			out = val
		}
	}

	return out
}

func MinIntInSlice(in []int) int {
	out := math.MaxInt
	for _, val := range in {
		if val < out {
			out = val
		}
	}

	return out
}
