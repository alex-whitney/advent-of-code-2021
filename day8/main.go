package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	patterns [][]string
	outputs  [][]string
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.patterns = make([][]string, len(in))
	d.outputs = make([][]string, len(in))
	for i, row := range in {
		parts := strings.Split(row, "|")
		d.patterns[i] = strings.Split(strings.TrimSpace(parts[0]), " ")
		d.outputs[i] = strings.Split(strings.TrimSpace(parts[1]), " ")
	}

	return nil
}

func intersect(a string, b string) string {
	out := ""
	for _, val := range a {
		if strings.Contains(b, string(val)) {
			out = out + string(val)
		}
	}

	return out
}

func toNum(in []string, code map[int]string) int {
	out := 0
	factor := 1000
	for _, val := range in {
		for k, v := range code {
			if len(val) == len(v) && intersect(val, v) == val {
				out = out + k*factor
			}
		}
		factor = factor / 10
	}
	return out
}

func decode(patterns []string) map[int]string {
	nums := make(map[int]string)

	// Collect uniques...
	for _, p := range patterns {
		if len(p) == 2 {
			// 1
			nums[1] = p
		} else if len(p) == 3 {
			// 7
			nums[7] = p
		} else if len(p) == 4 {
			// 4
			nums[4] = p
		} else if len(p) == 7 {
			nums[8] = p
		}
	}

	for _, p := range patterns {
		if len(p) == 5 {
			// 3 has the same chars as 1
			if intersect(nums[1], p) == nums[1] {
				nums[3] = p
			} else if len(intersect(nums[4], p)) == 3 {
				// 5 has the same chars as 4
				nums[5] = p
			} else {
				// this is a 2
				nums[2] = p
			}
		}
	}

	for _, p := range patterns {
		if len(p) == 6 {
			// 6 missing is part of 1
			if len(intersect(nums[1], p)) == 1 {
				nums[6] = p
			} else if len(intersect(nums[4], p)) == 3 {
				// find 0
				nums[0] = p
			} else {
				// this is a 9
				nums[9] = p
			}
		}
	}

	return nums
}

func (d *Today) Part1() (string, error) {
	counter := 0
	for _, p := range d.outputs {
		for _, val := range p {
			if len(val) == 2 {
				// 1
				counter++
			} else if len(val) == 3 {
				// 7
				counter++
			} else if len(val) == 4 {
				// 4
				counter++
			} else if len(val) == 7 {
				// 8
				counter++
			}
		}
	}

	return strconv.Itoa(counter), nil
}

func (d *Today) Part2() (string, error) {
	decode(d.patterns[0])
	sum := 0
	for i := range d.patterns {
		key := decode(d.patterns[i])
		sum += toNum(d.outputs[i], key)
	}

	return strconv.Itoa(sum), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
