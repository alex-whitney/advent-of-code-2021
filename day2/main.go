package main

import (
	"errors"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	instructions [][]string
}

func (d *Today) init() {
	lines, err := lib.ReadDelimitedFile("./input.txt", " ")
	if err != nil {
		panic(err)
	}
	d.instructions = lines
}

func (d *Today) Part1() (string, error) {
	x := 0
	y := 0
	for _, line := range d.instructions {
		num, err := strconv.Atoi(line[1])
		if err != nil {
			return "", err
		}

		switch line[0] {
		case "forward":
			x = x + num
		case "up":
			y = y - num
		case "down":
			y = y + num
		default:
			return "", errors.New("unknown direction: " + line[0])
		}
	}

	return strconv.Itoa(x * y), nil
}

func (d *Today) Part2() (string, error) {
	x := 0
	y := 0
	aim := 0
	for _, line := range d.instructions {
		num, err := strconv.Atoi(line[1])
		if err != nil {
			return "", err
		}

		switch line[0] {
		case "forward":
			x = x + num
			y = y + aim*num
		case "up":
			aim = aim - num
		case "down":
			aim = aim + num
		default:
			return "", errors.New("unknown direction: " + line[0])
		}
	}

	return strconv.Itoa(x * y), nil
}

func main() {
	day := &Today{}
	day.init()
	lib.Run(day)
}
