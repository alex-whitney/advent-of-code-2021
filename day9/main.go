package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	input [][]int
	lenX  int
	lenY  int
}

func (d *Today) Init(input string) error {
	rows, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.lenX = len(rows)
	d.lenY = len(rows[0])

	d.input = make([][]int, d.lenX)
	for r, row := range rows {
		d.input[r] = make([]int, d.lenY)
		for c, char := range row {
			d.input[r][c], err = strconv.Atoi(string(char))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Today) isLow(r int, c int) bool {
	val := d.input[r][c]
	isLow := true

	if r-1 >= 0 {
		isLow = isLow && d.input[r-1][c] > val
	}
	if r+1 < d.lenX {
		isLow = isLow && d.input[r+1][c] > val
	}
	if c-1 >= 0 {
		isLow = isLow && d.input[r][c-1] > val
	}
	if c+1 < d.lenY {
		isLow = isLow && d.input[r][c+1] > val
	}

	return isLow
}

func (d *Today) Part1() (string, error) {
	risk := 0

	for r, row := range d.input {
		for c, val := range row {
			if d.isLow(r, c) {
				risk += val + 1
			}
		}
	}

	return strconv.Itoa(risk), nil
}

func hash(p lib.Point) string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (d *Today) Part2() (string, error) {
	lowPoints := make([]lib.Point, 0)

	for r, row := range d.input {
		for c := range row {
			if d.isLow(r, c) {
				lowPoints = append(lowPoints, lib.Point{X: r, Y: c})
			}
		}
	}

	basins := make([]map[string]lib.Point, len(lowPoints))
	for i, center := range lowPoints {
		basins[i] = make(map[string]lib.Point)
		basin := basins[i]

		edges := []lib.Point{center}
		for len(edges) > 0 {
			newEdges := make([]lib.Point, 0)
			for _, edge := range edges {
				if _, ok := basin[hash(edge)]; ok {
					continue
				}
				if d.input[edge.X][edge.Y] == 9 {
					continue
				}
				basin[hash(edge)] = edge

				val := d.input[edge.X][edge.Y]
				r := edge.X
				c := edge.Y
				if r-1 >= 0 && d.input[r-1][c] > val {
					newEdges = append(newEdges, lib.Point{X: r - 1, Y: c})
				}
				if r+1 < d.lenX && d.input[r+1][c] > val {
					newEdges = append(newEdges, lib.Point{X: r + 1, Y: c})
				}
				if c-1 >= 0 && d.input[r][c-1] > val {
					newEdges = append(newEdges, lib.Point{X: r, Y: c - 1})
				}
				if c+1 < d.lenY && d.input[r][c+1] > val {
					newEdges = append(newEdges, lib.Point{X: r, Y: c + 1})
				}
			}

			edges = newEdges
		}
	}

	lens := make([]int, len(basins))
	for i, basin := range basins {
		lens[i] = len(basin)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lens)))

	return strconv.Itoa(lens[0] * lens[1] * lens[2]), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
