package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type board struct {
	values [boardSize][boardSize]int
	marks  [boardSize][boardSize]bool
}

func (board *board) mark(n int) {
	for row, r := range board.values {
		for col, val := range r {
			if val == n {
				board.marks[row][col] = true
			}
		}
	}
}

func (board *board) isBingo() bool {
	for row := range board.values {
		col := 0
		for col < boardSize && board.marks[row][col] {
			col++
		}
		if col == boardSize {
			return true
		}

		col = 0
		for col < boardSize && board.marks[col][row] {
			col++
		}
		if col == boardSize {
			return true
		}
	}
	return false
}

func (board *board) sumUnmarked() (result int) {
	for row, r := range board.values {
		for col, val := range r {
			if !board.marks[row][col] {
				result += val
			}
		}
	}
	return
}

func bothParts(boards []board, numbers []int) (int, int) {
	bingoCount, firstBingo, lastBingo := 0, 0, 0
	for _, n := range numbers {
		for i := range boards {
			board := &boards[i]
			if !board.isBingo() {
				board.mark(n)
				if board.isBingo() {
					bingoCount++
					if bingoCount == 1 {
						firstBingo = board.sumUnmarked() * n
					} else if bingoCount == len(boards) {
						lastBingo = board.sumUnmarked() * n
					}
				}
			}
		}
	}
	return firstBingo, lastBingo
}

func main() {
	boards, numbers := parseInput("day04/input.txt")
	part1, part2 := bothParts(boards, numbers)

	fmt.Println("Part 1 answer:", part1)
	fmt.Println("Part 2 answer:", part2)
}

func parseInput(filename string) ([]board, []int) {
	file, err := os.Open(filename)
	check(err)
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	scanner.Scan()
	for _, val := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(val)
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	var values [boardSize][boardSize]int
	var boards []board
	for scanner.Scan() {
		for row := 0; row < boardSize; row++ {
			scanner.Scan()
			for col, val := range strings.Fields(scanner.Text()) {
				n, err := strconv.Atoi(val)
				if err == nil {
					values[row][col] = n
				}
			}
		}
		boards = append(boards, board{values: values})
	}
	return boards, numbers
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
