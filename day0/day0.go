package main

import (
	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
}

func (d *Today) Part1() (string, error) {
	return "Hello", nil
}

func (d *Today) Part2() (string, error) {
	return "World", nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
