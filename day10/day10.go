package main

import (
	"os"
)

func main() {
	raw_input, err := os.ReadFile("day10.in")
	if err != nil {
		panic(err)
	}
	input := make([][]byte, 0)
	{
		s := 0
		for i, c := range raw_input {
			if c == '\n' {
				input = append(input, raw_input[s:i])
				s = i + 1
			}
		}

		for y := range input {
			for x := range input[y] {
				input[y][x] -= '0'
			}
		}
	}

	//fmt.Printf("%v\n", input)
	part1(input)
	part2(input)
}

func part1(input [][]byte) {
	ans := 0

	reached := make([][]bool, len(input))
	for y := range reached {
		reached[y] = make([]bool, len(input[y]))
	}
	clear_reached := func() {
		for y := range reached {
			for x := range reached[y] {
				reached[y][x] = false
			}
		}
	}

	var score_head func(val, y, x int) int
	score_head = func(val, y, x int) int {
		if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) {
			return 0
		}
		if input[y][x] != byte(val) {
			return 0
		}
		if val == 9 {
			if reached[y][x] {
				return 0
			}
			reached[y][x] = true
			return 1
		}

		return score_head(val+1, y-1, x) + score_head(val+1, y+1, x) + score_head(val+1, y, x-1) + score_head(val+1, y, x+1)
	}

	for y := range input {
		for x := range input[y] {
			if input[y][x] == 0 {
				clear_reached()
				ans += score_head(0, y, x)
			}
		}
	}

	println("Part 1:", ans)
}

func part2(input [][]byte) {
	ans := 0

	var score_head func(val, y, x int) int
	score_head = func(val, y, x int) int {
		if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) {
			return 0
		}
		if input[y][x] != byte(val) {
			return 0
		}
		if val == 9 {
			return 1
		}

		return score_head(val+1, y-1, x) + score_head(val+1, y+1, x) + score_head(val+1, y, x-1) + score_head(val+1, y, x+1)
	}

	for y := range input {
		for x := range input[y] {
			if input[y][x] == 0 {
				ans += score_head(0, y, x)
			}
		}
	}

	println("Part 2:", ans)
}
