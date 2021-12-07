package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	positions []int
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	d.positions, err = lib.ParseIntegerSlice(in, ",")
	if err != nil {
		return err
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	minCost := math.MaxInt
	minPos := -1

	for _, pos := range d.positions {
		cost := 0

		for _, val := range d.positions {
			cost += lib.AbsInt(pos - val)
		}

		if cost < minCost {
			minCost = cost
			minPos = pos
		}
	}

	fmt.Printf("pos: %d, fuel: %d\n", minPos, minCost)

	return strconv.Itoa(minCost), nil
}

func calcCost(a int, b int) int {
	diff := lib.AbsInt(a - b)

	if diff == 0 {
		return 0
	}

	// sum of numbers in a series
	return (diff * (diff + 1)) / 2
}

func (d *Today) Part2() (string, error) {
	minCost := math.MaxInt
	minPos := -1

	maxPos := lib.MaxInInSlice(d.positions)

	for i := 0; i <= maxPos; i++ {
		cost := 0

		for _, val := range d.positions {
			cost += calcCost(i, val)
		}

		if cost < minCost {
			minCost = cost
			minPos = i
		}
	}

	fmt.Printf("pos: %d, fuel: %d\n", minPos, minCost)

	return strconv.Itoa(minCost), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
