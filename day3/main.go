package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

// Probably should have just converted the binary strings to numbers and
// used bit ops, rather than string comparisons/conversions and conditionals

func countBits(in []string) []int {
	bits := make([]int, len(in[0]))
	for _, row := range in {
		for i, v := range row {
			if v == '0' {
				bits[i]--
			} else if v == '1' {
				bits[i]++
			}
		}
	}
	return bits
}

func filter(in []string, pos int, test rune) []string {
	out := make([]string, 0)
	for _, row := range in {
		if rune(row[pos]) == test {
			out = append(out, row)
		}
	}
	return out
}

type Today struct {
	input []string
}

func (d *Today) init() {
	in, err := lib.ReadStringFile("/input.txt")
	if err != nil {
		panic(err)
	}
	d.input = in
}

func (d *Today) Part1() (string, error) {
	bits := countBits(d.input)
	val1 := 0
	val2 := 0
	for _, bit := range bits {
		if bit > 0 {
			val1 = val1<<1 + 1
			val2 = val2 << 1
		} else {
			val1 = val1 << 1
			val2 = val2<<1 + 1
		}
	}
	fmt.Printf("val1: %d, val2: %d\n", val1, val2)
	return strconv.Itoa(val1 * val2), nil
}

func (d *Today) Part2() (string, error) {
	o2 := make([]string, len(d.input))
	copy(o2, d.input)
	co2 := make([]string, len(d.input))
	copy(co2, d.input)
	for pos := range d.input[0] {
		if len(o2) > 1 {
			bits := countBits(o2)
			bitVal := bits[pos]

			test := '1'
			if bitVal < 0 {
				test = '0'
			}
			o2 = filter(o2, pos, test)
		}
		if len(co2) > 1 {
			bits := countBits(co2)
			bitVal := bits[pos]

			test := '0'
			if bitVal < 0 {
				test = '1'
			}
			co2 = filter(co2, pos, test)
		}
	}

	if len(o2) != 1 {
		return "", errors.New("more than 1 item left for o2")
	}
	if len(co2) != 1 {
		return "", errors.New("more than 1 item left for co2")
	}

	v1, _ := strconv.ParseInt(o2[0], 2, 64)
	v2, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Printf("o2: %d, co2: %d\n", v1, v2)

	return strconv.FormatInt(v1*v2, 10), nil
}

func main() {
	day := &Today{}
	day.init()
	lib.Run(day)
}
