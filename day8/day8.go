package main

import (
	"os"
)

func main() {
	raw_input, err := os.ReadFile("day8.in")
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
	}

	part1(input)
	part2(input)
}

type Point struct {
	y int
	x int
}

func part1(input [][]byte) {
	h, w := len(input), len(input[0])

	antennas := make(map[byte][]Point)

	for y := range h {
		for x := range w {
			c := input[y][x]
			if c != '.' {
				a := antennas[c]
				a = append(a, Point{y, x})
				antennas[c] = a
			}
		}
	}

	antinodes := make([][]bool, h)
	for r := range antinodes {
		antinodes[r] = make([]bool, w)
	}

	ans := 0
	add_antinode := func(y int, x int) {
		if y < 0 || y >= h || x < 0 || x >= w {
			return
		}
		if antinodes[y][x] {
			return
		}
		antinodes[y][x] = true
		ans += 1
	}
	for _, points := range antennas {
		for i, p1 := range points {
			for _, p2 := range points[i+1:] {
				dy, dx := p1.y-p2.y, p1.x-p2.x

				add_antinode(p1.y+dy, p1.x+dx)
				add_antinode(p2.y-dy, p2.x-dx)
			}
		}
	}

	println("Part 1:", ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func part2(input [][]byte) {
	h, w := len(input), len(input[0])

	antennas := make(map[byte][]Point)

	for y := range h {
		for x := range w {
			c := input[y][x]
			if c != '.' {
				a := antennas[c]
				a = append(a, Point{y, x})
				antennas[c] = a
			}
		}
	}

	antinodes := make([][]bool, h)
	for r := range antinodes {
		antinodes[r] = make([]bool, w)
	}

	ans := 0
	add_antinode := func(y int, x int) {
		if y < 0 || y >= h || x < 0 || x >= w {
			return
		}
		if antinodes[y][x] {
			return
		}
		antinodes[y][x] = true
		ans += 1
	}
	for _, points := range antennas {
		for i, p1 := range points {
			for _, p2 := range points[i+1:] {
				dy, dx := p1.y-p2.y, p1.x-p2.x

				d := gcd(dy, dx)
				dy, dx = dy/d, dx/d

				for i := range 200 {
					i -= 100
					add_antinode(p1.y+dy*i, p1.x+dx*i)
				}
			}
		}
	}

	println("Part 1:", ans)
}
