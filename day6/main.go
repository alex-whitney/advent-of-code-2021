package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	fish []int
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadDelimitedFile(input, ",")
	if err != nil {
		return err
	}

	d.fish = make([]int, len(in[0]))
	for i, val := range in[0] {
		d.fish[i], _ = strconv.Atoi(val)
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	target := 80

	currentFish := make([]int, len(d.fish))
	copy(currentFish, d.fish)

	for day := 1; day <= target; day++ {
		newFish := make([]int, 0)

		for _, fish := range currentFish {
			if fish == 0 {
				newFish = append(newFish, 6, 8)
			} else {
				newFish = append(newFish, fish-1)
			}
		}

		currentFish = newFish
	}

	return strconv.Itoa(len(currentFish)), nil
}

func (d *Today) Part2() (string, error) {
	target := 256

	currentFish := make(map[int]int)
	for _, val := range d.fish {
		currentFish[val]++
	}

	for day := 1; day <= target; day++ {
		newFish := make(map[int]int)

		for cntr, num := range currentFish {
			if cntr == 0 {
				newFish[6] = newFish[6] + num
				newFish[8] = num
			} else {
				newFish[cntr-1] = newFish[cntr-1] + num
			}
		}

		currentFish = newFish
	}

	total := 0
	for _, v := range currentFish {
		total = total + v
	}

	return strconv.Itoa(total), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
