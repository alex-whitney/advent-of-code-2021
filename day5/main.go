package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type point struct {
	x int
	y int
}

func (p point) equal(p2 point) bool {
	return p.x == p2.x && p.y == p2.y
}

func (p point) copy() point {
	return point{
		x: p.x,
		y: p.y,
	}
}

func newPoint(in string) point {
	parts := strings.Split(in, ",")

	out := point{}
	out.x, _ = strconv.Atoi(strings.Trim(parts[0], " "))
	out.y, _ = strconv.Atoi(strings.Trim(parts[1], " "))

	return out
}

type line struct {
	start point
	end   point
}

type Today struct {
	lines []line

	maxX int
	maxY int
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.lines = make([]line, len(raw))
	for i, val := range raw {
		parts := strings.Split(val, "->")

		d.lines[i] = line{
			start: newPoint(parts[0]),
			end:   newPoint(parts[1]),
		}
	}

	for _, line := range d.lines {
		if line.start.x > d.maxX {
			d.maxX = line.start.x
		}
		if line.start.y > d.maxY {
			d.maxY = line.start.y
		}
		if line.end.x > d.maxX {
			d.maxX = line.end.x
		}
		if line.end.y > d.maxY {
			d.maxY = line.end.y
		}
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	chart := make([][]int, d.maxX+1)
	for i := range chart {
		chart[i] = make([]int, d.maxY+1)
	}

	for _, line := range d.lines {
		if line.start.x == line.end.x {
			start := line.start.copy()
			end := line.end.copy()
			if line.end.y < line.start.y {
				c := start
				start = end
				end = c
			}
			for start.y <= end.y {
				chart[start.x][start.y]++
				start.y++
			}
		} else if line.start.y == line.end.y {
			start := line.start.copy()
			end := line.end.copy()
			if line.end.x < line.start.x {
				c := start
				start = end
				end = c
			}
			for start.x <= end.x {
				chart[start.x][start.y]++
				start.x++
			}
		}
	}

	count := 0
	for _, row := range chart {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d *Today) Part2() (string, error) {
	chart := make([][]int, d.maxX+1)
	for i := range chart {
		chart[i] = make([]int, d.maxY+1)
	}

	for _, line := range d.lines {
		start := line.start.copy()
		end := line.end.copy()

		incX := 0
		if line.start.x < line.end.x {
			incX = 1
		} else if line.start.x > line.end.x {
			incX = -1
		}

		incY := 0
		if line.start.y < line.end.y {
			incY = 1
		} else if line.start.y > line.end.y {
			incY = -1
		}

		for {
			chart[start.x][start.y]++

			if start.equal(end) {
				break
			}

			start.x += incX
			start.y += incY
		}
	}

	count := 0
	for _, row := range chart {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
