package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	input lib.IntMatrix
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.input, err = lib.NewIntMatrix(in, "")
	if err != nil {
		return err
	}

	return nil
}

func inc(state [][]int, r int, c int) bool {
	state[r][c]++
	return state[r][c] == 10
}

func reset(state [][]int) {
	for r, row := range state {
		for c, val := range row {
			if val > 9 {
				state[r][c] = 0
			}
		}
	}
}

func (d *Today) Part1() (string, error) {
	state := d.input.Copy()

	flashes := 0
	for step := 0; step < 100; step++ {
		newFlashes := make(map[string]*lib.Point)
		handledFlashes := make(map[string]*lib.Point)

		for r, row := range state {
			for c := range row {
				if inc(state, r, c) {
					p := lib.NewPoint(r, c)
					newFlashes[p.Hash()] = p
				}
			}
		}

		for len(newFlashes) > 0 {
			newFlashes2 := make(map[string]*lib.Point)

			for _, p := range newFlashes {
				if _, ok := handledFlashes[p.Hash()]; ok {
					continue
				}

				for _, dr := range []int{-1, 0, 1} {
					for _, dc := range []int{-1, 0, 1} {
						if dr == 0 && dc == 0 {
							continue
						}

						r := p.X + dr
						c := p.Y + dc
						if state.InBounds(r, c) && inc(state, r, c) {
							p := lib.NewPoint(r, c)
							newFlashes2[p.Hash()] = p
						}
					}
				}
				handledFlashes[p.Hash()] = p
			}
			newFlashes = newFlashes2
		}

		reset(state)
		flashes += len(handledFlashes)
	}

	return strconv.Itoa(flashes), nil
}

func (d *Today) Part2() (string, error) {
	state := d.input.Copy()

	allFlashed := false
	counter := 0
	for !allFlashed {
		newFlashes := make(map[string]*lib.Point)
		handledFlashes := make(map[string]*lib.Point)

		for r, row := range state {
			for c := range row {
				if inc(state, r, c) {
					p := lib.NewPoint(r, c)
					newFlashes[p.Hash()] = p
				}
			}
		}

		for len(newFlashes) > 0 {
			newFlashes2 := make(map[string]*lib.Point)

			for _, p := range newFlashes {
				if _, ok := handledFlashes[p.Hash()]; ok {
					continue
				}

				for _, dr := range []int{-1, 0, 1} {
					for _, dc := range []int{-1, 0, 1} {
						if dr == 0 && dc == 0 {
							continue
						}

						r := p.X + dr
						c := p.Y + dc
						if state.InBounds(r, c) && inc(state, r, c) {
							p := lib.NewPoint(r, c)
							newFlashes2[p.Hash()] = p
						}
					}
				}
				handledFlashes[p.Hash()] = p
			}
			newFlashes = newFlashes2
		}

		reset(state)
		allFlashed = len(handledFlashes) == 100
		counter++
	}

	return strconv.Itoa(counter), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
