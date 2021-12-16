package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bothParts(filename string) (int, int) {
	file, err := os.Open(filename)
	check(err)
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var aim int
	var pos int
	var depth int

	for scanner.Scan() {
		command := scanner.Text()
		scanner.Scan()
		value, err := strconv.Atoi(scanner.Text())
		check(err)

		switch command {
		case "forward":
			pos += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			fmt.Fprintln(os.Stderr, "wrong input format!")
			os.Exit(1)
		}
	}
	return pos * aim, pos * depth
}

func main() {
	part1, part2 := bothParts("day02/input.txt")

	fmt.Println("Part 1 answer:", part1)
	fmt.Println("Part 2 answer:", part2)
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
