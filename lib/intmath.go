package lib

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) Hash() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

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
