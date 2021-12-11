package lib

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
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

type IntMatrix [][]int

func NewIntMatrix(input []string, delimiter string) (IntMatrix, error) {
	out := make(IntMatrix, len(input))
	lenC := 0
	for r, row := range input {
		parts := strings.Split(row, delimiter)
		out[r] = make([]int, len(parts))

		if r == 0 {
			lenC = len(parts)
		} else if len(parts) != lenC {
			return nil, errors.New("inconsistent column size")
		}

		for c, val := range parts {
			var err error
			out[r][c], err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
		}
	}

	return out, nil
}

func (m IntMatrix) Copy() IntMatrix {
	out := make(IntMatrix, len(m))
	for r, row := range m {
		out[r] = make([]int, len(row))
		for c, val := range row {
			out[r][c] = val
		}
	}
	return out
}

func (m IntMatrix) InBounds(r int, c int) bool {
	return r >= 0 && c >= 0 && r < len(m) && c < len(m[0])
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
