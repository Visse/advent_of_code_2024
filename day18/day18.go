package main

import (
	"fmt"
	"io"
	"os"
)

type P struct {
	y, x int
}

func main() {
	input := func() []P {
		file, err := os.Open("day18.in")
		if err != nil {
			panic(err)
		}
		input := make([]P, 0)
		for {
			var p P
			_, err = fmt.Fscanf(file, "%v,%v\n", &p.x, &p.y)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			input = append(input, p)
		}
		return input
	}()
	part1(input)
	part2(input)
}

func part1(input []P) {
	visited := make(map[P]bool)
	const UNREACHABLE int = 1000000000
	//var step func(p P, d int) int

	for _, p := range input[:1024] {
		visited[p] = true
	}

	type S struct {
		p P
		d int
	}
	stack := make([]S, 0)

	swap := func(a, b int) {
		t := stack[a]
		stack[a] = stack[b]
		stack[b] = t
	}

	add := func(s S) {
		i := len(stack)
		stack = append(stack, s)
		for i > 0 {
			p := (i - 1) / 2
			if stack[p].d < stack[i].d {
				break
			}
			swap(i, p)
			i = p
		}
	}
	pop := func() S {
		s := stack[0]
		stack[0] = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i := 0
		for {
			if len(stack) > (i*2+1) && stack[i].d > stack[i*2+1].d {
				swap(i, i*2+1)
				i = i*2 + 1
			} else if len(stack) > (i*2+2) && stack[i].d > stack[i*2+2].d {
				swap(i, i*2+2)
				i = i*2 + 2
			} else {
				break
			}
		}
		return s
	}

	add(S{P{0, 0}, 0})

	ans := 0
	for {
		if len(stack) == 0 {
			return
		}
		s := pop()
		p, d := s.p, s.d
		if p.y < 0 || p.x < 0 || p.y > 70 || p.x > 70 {
			continue
		}
		if visited[p] {
			continue
		}
		visited[p] = true
		if s.p == (P{y: 70, x: 70}) {
			ans = s.d
			break
		}
		add(S{P{y: p.y - 1, x: p.x}, d + 1})
		add(S{P{y: p.y + 1, x: p.x}, d + 1})
		add(S{P{y: p.y, x: p.x - 1}, d + 1})
		add(S{P{y: p.y, x: p.x + 1}, d + 1})
	}
	println("Part 1:", ans)
}

func part2(input []P) {
	l, u := 1024, len(input)
	for u-l > 1 {
		m := (l + u) / 2

		visited := make(map[P]bool)
		const UNREACHABLE int = 1000000000
		//var step func(p P, d int) int

		for _, p := range input[:m] {
			visited[p] = true
		}

		type S struct {
			p P
			d int
		}
		stack := make([]S, 0)

		swap := func(a, b int) {
			t := stack[a]
			stack[a] = stack[b]
			stack[b] = t
		}

		add := func(s S) {
			i := len(stack)
			stack = append(stack, s)
			for i > 0 {
				p := (i - 1) / 2
				if stack[p].d < stack[i].d {
					break
				}
				swap(i, p)
				i = p
			}
		}
		pop := func() S {
			s := stack[0]
			stack[0] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i := 0
			for {
				if len(stack) > (i*2+1) && stack[i].d > stack[i*2+1].d {
					swap(i, i*2+1)
					i = i*2 + 1
				} else if len(stack) > (i*2+2) && stack[i].d > stack[i*2+2].d {
					swap(i, i*2+2)
					i = i*2 + 2
				} else {
					break
				}
			}
			return s
		}

		add(S{P{0, 0}, 0})
		reached := false
		for !reached {
			if len(stack) == 0 {
				break
			}
			s := pop()
			p, d := s.p, s.d
			if p.y < 0 || p.x < 0 || p.y > 70 || p.x > 70 {
				continue
			}
			if visited[p] {
				continue
			}
			visited[p] = true
			if s.p == (P{y: 70, x: 70}) {
				reached = true
				break
			}
			add(S{P{y: p.y - 1, x: p.x}, d + 1})
			add(S{P{y: p.y + 1, x: p.x}, d + 1})
			add(S{P{y: p.y, x: p.x - 1}, d + 1})
			add(S{P{y: p.y, x: p.x + 1}, d + 1})
		}

		if reached {
			l = m
		} else {
			u = m
		}
	}

	p := input[l]
	fmt.Printf("Part 2: %v,%v", p.x, p.y)
}
