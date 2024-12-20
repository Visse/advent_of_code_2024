package main

import (
	"fmt"
	"os"
)

type Robot struct {
	p_x, p_y int
	v_x, v_y int
}

func main() {

	input := func() []Robot {
		input := make([]Robot, 0)

		file, err := os.Open("day14.in")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for {
			var p_x, p_y int
			var v_x, v_y int
			_, err = fmt.Fscanf(file, "p=%v,%v v=%v,%v\n", &p_x, &p_y, &v_x, &v_y)
			if err != nil {
				break
			}

			input = append(input, Robot{p_x, p_y, v_x, v_y})
		}
		return input
	}()

	part1(input)
	part2(input)
}

func part1(input []Robot) {
	const W = 101
	const H = 103

	quads := make([]int, 4)

	for _, r := range input {
		x := r.p_x + r.v_x*100
		y := r.p_y + r.v_y*100

		x = ((x % W) + W) % W
		y = ((y % H) + H) % H

		if x == W/2 || y == H/2 {
			// center line
			continue
		}

		q := 0
		if x < W/2 {
			q += 2
		}
		if y < H/2 {
			q += 1
		}
		quads[q] += 1
	}

	ans := quads[0] * quads[1] * quads[2] * quads[3]
	println("Part 1:", ans)
}

func part2(input []Robot) {
	const W = 101
	const H = 103

	m := make([][]int, H)
	for y := range H {
		m[y] = make([]int, W)
	}

	for s := range 10000 {
		c := true
		for _, r := range input {
			x := r.p_x + r.v_x*s
			y := r.p_y + r.v_y*s

			x = ((x % W) + W) % W
			y = ((y % H) + H) % H

			if x == W/2 || y == H/2 {
				// center line
				continue
			}

			m[y][x] += 1
			if m[y][x] != 1 {
				c = false
				break
			}
		}

		for y := range H {
			for x := range W {
				if c {
					if m[y][x] == 0 {
						print(" ")
					} else {
						print(m[y][x])
					}
				}
				m[y][x] = 0
			}
			if c {
				println()
			}
		}
		if c {
			println("Ans:", s)
			break
		}
	}

}
