package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	input [][]int
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.input = make([][]int, len(in))
	for i, row := range in {
		d.input[i] = make([]int, len(row))
		for j, val := range row {
			d.input[i][j], err = strconv.Atoi(string(val))
			if err != nil {
				return err
			}
		}
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
	state := make([][]int, len(d.input))
	for r, row := range d.input {
		state[r] = make([]int, len(row))
		for c, val := range row {
			state[r][c] = val
		}
	}

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
						if r >= 0 && r < 10 && c >= 0 && c < 10 {
							if inc(state, r, c) {
								p := lib.NewPoint(r, c)
								newFlashes2[p.Hash()] = p
							}
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
	state := make([][]int, len(d.input))
	for r, row := range d.input {
		state[r] = make([]int, len(row))
		for c, val := range row {
			state[r][c] = val
		}
	}

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
						if r >= 0 && r < 10 && c >= 0 && c < 10 {
							if inc(state, r, c) {
								p := lib.NewPoint(r, c)
								newFlashes2[p.Hash()] = p
							}
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
