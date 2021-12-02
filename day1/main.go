package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	nums []int
}

func (d *Today) init() {
	input, err := lib.ReadIntegerFile("/input.txt")
	if err != nil {
		panic(err)
	}

	d.nums = input
}

func (d *Today) Part1() (string, error) {
	count := 0
	for i := 1; i < len(d.nums); i++ {
		if d.nums[i] > d.nums[i-1] {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func (d *Today) Part2() (string, error) {
	count := 0
	for i := 3; i < len(d.nums); i++ {
		if d.nums[i] > d.nums[i-3] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func main() {
	day := &Today{}
	day.init()
	lib.Run(day)
}
