package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type stack []rune

func (s stack) push(v rune) stack {
	return append(s, v)
}

func (s stack) pop() (stack, rune) {
	return s[:len(s)-1], s[len(s)-1]
}

func (s stack) peek() rune {
	return s[len(s)-1]
}

func getOpenToClose() map[rune]rune {
	return map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
}

type Today struct {
	input []string
}

func (d *Today) Init(input string) error {
	var err error
	if d.input, err = lib.ReadStringFile(input); err != nil {
		return err
	}
	return nil
}

func (d *Today) Part1() (string, error) {
	score := 0
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	openToClose := getOpenToClose()

	for _, row := range d.input {
		open := make(stack, 0)
		for _, char := range row {
			if strings.ContainsRune("{([<", char) {
				open = open.push(openToClose[char])
			} else {
				if open.peek() == char {
					open, _ = open.pop()
				} else {
					score += points[char]
					break
				}
			}
		}
	}

	return strconv.Itoa(score), nil
}

func (d *Today) Part2() (string, error) {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	openToClose := getOpenToClose()

	scores := make([]int, 0)
	for _, row := range d.input {
		open := make(stack, 0)
		for _, char := range row {
			if strings.ContainsRune("{([<", char) {
				open = open.push(openToClose[char])
			} else {
				if open.peek() == char {
					open, _ = open.pop()
				} else {
					open = make(stack, 0)
					break
				}
			}
		}

		if len(open) == 0 {
			continue
		}

		score := 0
		for len(open) > 0 {
			var char rune
			open, char = open.pop()

			score = score*5 + points[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return strconv.Itoa(scores[(len(scores)-1)/2]), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
