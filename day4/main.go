package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Today struct {
	numbers []int
	boards  [][][]int
}

func (d *Today) Init(input string) error {
	contents, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	parts := strings.Split(contents, "\n\n")

	strNum := strings.Split(parts[0], ",")
	d.numbers = make([]int, len(strNum))
	for i, n := range strNum {
		d.numbers[i], _ = strconv.Atoi(n)
	}

	d.boards = make([][][]int, len(parts)-1)
	for i := 1; i < len(parts); i++ {
		rows := strings.Split(parts[i], "\n")
		d.boards[i-1] = make([][]int, len(rows))
		for r, row := range rows {
			cells := strings.Split(row, " ")
			d.boards[i-1][r] = make([]int, 0)
			for _, cell := range cells {
				if cell != "" {
					ival, _ := strconv.Atoi(cell)
					d.boards[i-1][r] = append(d.boards[i-1][r], ival)
				}

			}
		}
	}

	return nil
}

func checkBingo(board [][]int) bool {
	nRow := len(board)
	nCol := len(board[0])

	for _, row := range board {
		bingo := true
		for _, val := range row {
			bingo = bingo && (val < 0)
		}
		if bingo {
			return true
		}
	}

	for c := 0; c < nCol; c++ {
		bingo := true
		for r := 0; r < nRow; r++ {
			bingo = bingo && (board[r][c] < 0)
		}
		if bingo {
			return true
		}
	}

	return false
}

func markNumber(board [][]int, val int) {
	for r, row := range board {
		for c, v := range row {
			if val == v {
				board[r][c] = -1
			}
		}
	}
}

func (d *Today) Part1() (string, error) {
	var winningBoard [][]int

	c := 0
	for winningBoard == nil {
		nextNumber := d.numbers[c]
		c++

		for _, board := range d.boards {
			markNumber(board, nextNumber)
			if checkBingo(board) {
				winningBoard = board
			}
		}
	}

	score := 0
	for _, row := range winningBoard {
		for _, v := range row {
			if v > 0 {
				score += v
			}
		}
	}

	return strconv.Itoa(score * d.numbers[c-1]), nil
}

func (d *Today) Part2() (string, error) {
	// if running both part1 and part2, part1 mutates the boards -- but this is OK
	boards := d.boards

	c := 0
	for len(boards) > 1 {
		nextNumber := d.numbers[c]
		c++

		nextBoards := make([][][]int, 0)
		for _, board := range boards {
			markNumber(board, nextNumber)
			if !checkBingo(board) {
				nextBoards = append(nextBoards, board)
			}
		}

		boards = nextBoards
	}

	// tricky. Missed this originally, since the sample problem didn't quite need it!
	for !checkBingo(boards[0]) {
		nextNumber := d.numbers[c]
		c++

		markNumber(boards[0], nextNumber)
	}

	score := 0
	for _, row := range boards[0] {
		for _, v := range row {
			if v > 0 {
				score += v
			}
		}
	}

	return strconv.Itoa(score * d.numbers[c-1]), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
