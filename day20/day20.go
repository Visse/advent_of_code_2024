package main

import (
	"os"
)

func main() {
	input := func() [][]byte {
		file, err := os.ReadFile("day20.in")
		if err != nil {
			panic(err)
		}
		input := make([][]byte, 0)
		for len(file) != 0 {
			i := 0
			for ; i < len(file) && file[i] != '\n'; i++ {

			}
			if i != 0 {
				input = append(input, file[:i])
			}
			file = file[i+1:]
		}

		return input
	}()

	part1(input)
	part2(input)
}

type P struct {
	y, x int
}
type S struct {
	p P
	d int
}

func push_heap(heap *[]S, v S) {
	h := *heap
	h = append(h, v)
	*heap = h

	i := len(h) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h[p].d < h[i].d {
			break
		}

		t := h[p]
		h[p] = h[i]
		h[i] = t
		i = p
	}
}

func pop_heap(heap *[]S) S {
	h := *heap
	v := h[0]
	h[0] = h[len(h)-1]
	*heap = h[:len(h)-1]

	i := 0
	for {
		c1 := i*2 + 1
		c2 := i*2 + 2

		l := c1
		if c2 < len(h) && h[c2].d < h[c1].d {
			l = c2
		}
		if l < len(h) && h[l].d < h[i].d {
			t := h[i]
			h[i] = h[l]
			h[l] = t
			i = l
		} else {
			break
		}
	}

	return v
}

func part1(input [][]byte) {
	var s, e P
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				s = P{y, x}
			} else if input[y][x] == 'E' {
				e = P{y, x}
			}
		}
	}

	flood_fill := func(p P) [][]int {
		d := make([][]int, len(input))
		for y := range input {
			d[y] = make([]int, len(input[y]))
			for x := range input[y] {
				d[y][x] = -1
			}
		}
		q := make([]S, 0)
		push_heap(&q, S{p, 0})

		for len(q) != 0 {
			s := pop_heap(&q)
			y, x := s.p.y, s.p.x
			if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) {
				continue
			}
			if d[y][x] != -1 {
				continue
			}

			if input[y][x] != '#' {
				d[y][x] = s.d
				push_heap(&q, S{P{y - 1, x}, s.d + 1})
				push_heap(&q, S{P{y + 1, x}, s.d + 1})
				push_heap(&q, S{P{y, x - 1}, s.d + 1})
				push_heap(&q, S{P{y, x + 1}, s.d + 1})
			}
		}

		return d
	}

	dist_S := flood_fill(s)
	dist_E := flood_fill(e)

	d := dist_S[e.y][e.x]

	ans := 0
	for y := range input {
		for x := range input[y] {
			for dy := range 7 {
				dy = dy - 3
				for dx := range 7 {
					dx = dx - 3

					t := 0
					if dy < 0 {
						t -= dy
					} else {
						t += dy
					}
					if dx < 0 {
						t -= dx
					} else {
						t += dx
					}
					if t > 2 {
						continue
					}

					ey, ex := y+dy, x+dx
					if ey < 0 || ey >= len(input) || ex < 0 || ex >= len(input[ey]) {
						continue
					}
					if input[y][x] == '#' || input[ey][ex] == '#' {
						continue
					}
					c := dist_E[ey][ex] + dist_S[y][x] + t
					if c+99 < d {
						ans++
					}

				}
			}
		}
	}

	println("Part 1:", ans)
}

func part2(input [][]byte) {
	var s, e P
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				s = P{y, x}
			} else if input[y][x] == 'E' {
				e = P{y, x}
			}
		}
	}

	flood_fill := func(p P) [][]int {
		d := make([][]int, len(input))
		for y := range input {
			d[y] = make([]int, len(input[y]))
			for x := range input[y] {
				d[y][x] = -1
			}
		}
		q := make([]S, 0)
		push_heap(&q, S{p, 0})

		for len(q) != 0 {
			s := pop_heap(&q)
			y, x := s.p.y, s.p.x
			if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) {
				continue
			}
			if d[y][x] != -1 {
				continue
			}

			if input[y][x] != '#' {
				d[y][x] = s.d
				push_heap(&q, S{P{y - 1, x}, s.d + 1})
				push_heap(&q, S{P{y + 1, x}, s.d + 1})
				push_heap(&q, S{P{y, x - 1}, s.d + 1})
				push_heap(&q, S{P{y, x + 1}, s.d + 1})
			}
		}

		return d
	}

	dist_S := flood_fill(s)
	dist_E := flood_fill(e)

	d := dist_S[e.y][e.x]

	ans := 0
	for y := range input {
		for x := range input[y] {
			for dy := range 41 {
				dy = dy - 20
				for dx := range 41 {
					dx = dx - 20

					t := 0
					if dy < 0 {
						t -= dy
					} else {
						t += dy
					}
					if dx < 0 {
						t -= dx
					} else {
						t += dx
					}
					if t > 20 {
						continue
					}

					ey, ex := y+dy, x+dx
					if ey < 0 || ey >= len(input) || ex < 0 || ex >= len(input[ey]) {
						continue
					}
					if input[y][x] == '#' || input[ey][ex] == '#' {
						continue
					}
					c := dist_E[ey][ex] + dist_S[y][x] + t
					if c+99 < d {
						ans++
					}

				}
			}
		}
	}

	println("Part 2:", ans)
}
