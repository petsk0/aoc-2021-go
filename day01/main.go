package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var result []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		result = append(result, number)
	}
	return result
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func part1(input []int) (res int) {
	if len(input) < 2 {
		return
	}

	for i, n := range input[:len(input)-1] {
		if n < input[i+1] {
			res++
		}
	}
	return
}

func part2(input []int) int {
	if len(input) < 3 {
		return 0
	}

	for i := range input[:len(input)-2] {
		input[i] = input[i] + input[i+1] + input[i+2]
	}
	return part1(input[:len(input)-2])
}

func main() {
	input := parseInput("day01/input.txt")

	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}
