package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(data [][]byte) int {
	epsilon, gamma := 0, 0
	ones := make([]int, len(data[0]))
	for _, row := range data {
		for col, val := range row {
			ones[col] += int(val)
		}
	}

	for i, count := range ones {
		shift := len(ones) - i - 1
		if count < len(data)-count {
			epsilon |= 1 << shift
		} else {
			gamma |= 1 << shift
		}
	}
	return epsilon * gamma
}

func solve(data [][]byte, oxygen bool) int {
	removed := make([]bool, len(data))
	ones := make([]int, len(data[0]))
	for col := range data[0] {
		remaining := 0
		for row, val := range data {
			if !removed[row] {
				ones[col] += int(val[col])
				remaining++
			}
		}
		if remaining == 1 {
			break
		}
		moreOnes := ones[col] >= (remaining - ones[col])
		for row := range data {
			if oxygen {
				if !removed[row] {
					val := data[row][col]
					if moreOnes && val == 0 || !moreOnes && val == 1 {
						removed[row] = true
					}
				}
			} else {
				if !removed[row] {
					val := data[row][col]
					if moreOnes && val == 1 || !moreOnes && val == 0 {
						removed[row] = true
					}
				}
			}
		}
	}

	for row := range data {
		if !removed[row] {
			result := 0
			for col, val := range data[row] {
				if val == 1 {
					result |= 1 << (len(data[row]) - col - 1)
				}
			}
			return result
		}
	}
	return 0
}

func part2(data [][]byte) int {
	oxygen := solve(data, true)
	co2 := solve(data, false)
	return oxygen * co2
}

func main() {
	input := parseInput("day03/input.txt")

	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func parseInput(filename string) [][]byte {
	file, err := os.Open(filename)
	check(err)
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data [][]byte
	for scanner.Scan() {
		tmp := scanner.Bytes()
		bytes := make([]byte, len(tmp))
		copy(bytes, tmp)
		for i, val := range bytes {
			bytes[i] = val - 48
		}
		data = append(data, bytes)
	}
	return data
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
