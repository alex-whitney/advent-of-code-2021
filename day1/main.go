package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	nums []int
}

func (d *Today) Init(file string) error {
	input, err := lib.ReadIntegerFile(file)
	if err != nil {
		return err
	}

	d.nums = input
	return nil
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
	lib.Run(day)
}
