package main

import (
	"os"
)

func main() {
	raw_input, err := os.ReadFile("day12.in")
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

	//fmt.Printf("%v\n", input)
	part1(input)
	part2(input)
}

func part1(input [][]byte) {
	region := make([][]int, len(input))

	for y := range input {
		region[y] = make([]int, len(input[y]))
	}

	type P struct {
		y, x int
	}
	stack := make([]P, 0)

	region_id := 1

	ans := 0
	for y := range input {
		for x := range input[y] {
			if region[y][x] != 0 {
				continue
			}

			c := input[y][x]

			perimeter := 0
			size := 0

			stack = append(stack, P{y, x})
			for {
				if len(stack) == 0 {
					break
				}
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if p.y < 0 || p.y >= len(input) || p.x < 0 || p.x >= len(input[y]) {
					perimeter += 1
					continue
				}
				if c != input[p.y][p.x] {
					perimeter += 1
					continue
				}
				if region[p.y][p.x] == region_id {
					continue
				}
				region[p.y][p.x] = region_id
				size += 1

				stack = append(stack, P{p.y - 1, p.x})
				stack = append(stack, P{p.y + 1, p.x})
				stack = append(stack, P{p.y, p.x - 1})
				stack = append(stack, P{p.y, p.x + 1})
			}

			//fmt.Printf("%vx%v[%v] %c - %v:%v\n", y, x, region_id, c, perimeter, size)
			ans += perimeter * size
			region_id += 1
		}
	}

	println("Part 1", ans)
}

func part2(input [][]byte) {
	region := make([][]int, len(input))

	for y := range input {
		region[y] = make([]int, len(input[y]))
	}

	type P struct {
		y, x int
	}
	stack := make([]P, 0)

	region_size := make([]int, 1)
	region_edges := make([]int, 1)

	for y := range input {
		for x := range input[y] {
			if region[y][x] != 0 {
				continue
			}

			c := input[y][x]

			size := 0
			region_id := len(region_size)

			stack = append(stack, P{y, x})
			for {
				if len(stack) == 0 {
					break
				}
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if p.y < 0 || p.y >= len(input) || p.x < 0 || p.x >= len(input[y]) {
					continue
				}
				if c != input[p.y][p.x] {
					continue
				}
				if region[p.y][p.x] == region_id {
					continue
				}
				region[p.y][p.x] = region_id
				size += 1

				stack = append(stack, P{p.y - 1, p.x})
				stack = append(stack, P{p.y + 1, p.x})
				stack = append(stack, P{p.y, p.x - 1})
				stack = append(stack, P{p.y, p.x + 1})
			}

			//fmt.Printf("%vx%v[%v] %c - %v:%v\n", y, x, region_id, c, perimeter, size)
			region_size = append(region_size, size)
			region_edges = append(region_edges, 0)
		}
	}

	get_region := func(y, x int) int {
		if y < 0 || y >= len(region) {
			return 0
		}
		if x < 0 || x >= len(region[y]) {
			return 0
		}
		return region[y][x]
	}

	for y := 0; y <= len(input); y++ {
		for x := 0; x <= len(input[0]); x++ {
			//  0 1
			//  2 3
			r0, r1, r2, r3 := get_region(y-1, x-1), get_region(y-1, x), get_region(y, x-1), get_region(y, x)

			if r0 == r1 && r0 == r2 && r0 == r3 {
				continue
			}

			if (r0 != r1 && r0 != r2) || (r0 == r1 && r0 == r2) {
				region_edges[r0] += 1
			}
			if (r1 != r0 && r1 != r3) || (r1 == r0 && r1 == r3) {
				region_edges[r1] += 1
			}
			if (r2 != r0 && r2 != r3) || (r2 == r0 && r2 == r3) {
				region_edges[r2] += 1
			}
			if (r3 != r2 && r3 != r1) || (r3 == r2 && r3 == r1) {
				region_edges[r3] += 1
			}
		}
	}

	ans := 0
	for r := range region_size {
		if r == 0 {
			continue
		}

		ans += region_size[r] * region_edges[r]
	}
	println("Part 2", ans)
}
