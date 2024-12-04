package main

import (
	"os"
)

func main() {
	load_input := func() [][]byte {
		raw, err := os.ReadFile("day4.in")
		if err != nil {
			panic(err)
		}

		input := [][]byte{}
		s := 0
		for p, c := range raw {
			if c == '\n' {
				input = append(input, raw[s:p])
				s = p + 1
			}
		}
		return input
	}

	input := load_input()
	part1(input)
	part2(input)
}

func count_words(input [][]byte, y int, x int, d int, c int) int {
	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	leters := []byte{'X', 'M', 'A', 'S'}

	if y < 0 || y >= len(input) {
		return 0
	}
	line := input[y]
	if x < 0 || x >= len(line) {
		return 0
	}
	if line[x] != leters[c] {
		return 0
	}
	c += 1
	if c == len(leters) {
		//fmt.Printf("Found %c at %v %v %s\n", leters[c-1], y, x, string(line))
		return 1
	}

	total := 0

	total += count_words(input, y+dirs[d][0], x+dirs[d][1], d, c)
	if total != 0 {
		//fmt.Printf("At %v %v %c\n", y, x, leters[c-1])
		return 1
	}

	return total
}

func part1(input [][]byte) {
	total := 0

	for y, line := range input {
		for x, _ := range line {
			for d := 0; d < 8; d += 1 {
				total += count_words(input, y, x, d, 0)
			}
		}
	}

	println("Part 1:", total)
}

func part2(input [][]byte) {
	check := func(y int, x int, c byte) bool {
		if y < 0 || y >= len(input) {
			return false
		}
		line := input[y]
		if x < 0 || x >= len(line) {
			return false
		}
		return line[x] == c
	}
	check_diag := func(y int, x int, diag int) bool {
		if check(y+diag, x-1, 'M') && check(y-diag, x+1, 'S') {
			return true
		}
		if check(y+diag, x-1, 'S') && check(y-diag, x+1, 'M') {
			return true
		}
		return false
	}

	total := 0
	for y, line := range input {
		for x, c := range line {
			if c == 'A' {
				if check_diag(y, x, -1) && check_diag(y, x, 1) {
					total += 1
				}
			}
		}
	}

	println("Part 2:", total)
}
