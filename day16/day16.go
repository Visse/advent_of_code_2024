package main

import (
	"os"
	"sort"
)

func main() {
	input := func() [][]byte {
		file, err := os.ReadFile("day16.in")
		if err != nil {
			panic(err)
		}

		input := make([][]byte, 0)

		s := 0
		for i := range file {
			if file[i] == '\n' {
				input = append(input, file[s:i])
				s = i + 1
			}
		}
		return input
	}()

	part1(input)
	part2(input)
}

func part1(input [][]byte) {
	type P struct {
		y, x int
	}
	type S struct {
		p, d  P
		score int
	}

	state := make([]S, 0)
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				state = append(state, S{p: P{y, x}, d: P{0, 1}, score: 0})
			}
		}
	}

	visited := make(map[S]bool)
	best_score := -1
	for {
		if len(state) == 0 {
			break
		}

		s := state[len(state)-1]
		//fmt.Printf("%v\n", s)
		state = state[:len(state)-1]

		v := s
		v.score = 0
		if visited[v] || input[s.p.y][s.p.x] == '#' {
			continue
		}
		visited[v] = true

		if input[s.p.y][s.p.x] == 'E' {
			if best_score == -1 || s.score < best_score {
				best_score = s.score
				break
			}
		}

		state = append(state, S{
			p:     s.p,
			d:     P{y: -s.d.x, x: s.d.y},
			score: s.score + 1000,
		})
		state = append(state, S{
			p:     s.p,
			d:     P{y: s.d.x, x: -s.d.y},
			score: s.score + 1000,
		})
		state = append(state, S{
			p:     P{y: s.p.y + s.d.y, x: s.p.x + s.d.x},
			d:     s.d,
			score: s.score + 1,
		})

		sort.Slice(state, func(i, j int) bool {
			return state[i].score > state[j].score
		})
	}
	println("Part 1:", best_score)
}

func part2(input [][]byte) {
	type P struct {
		y, x int
	}
	type S struct {
		p    P
		d, s int
	}

	dir := func(d int) P {
		switch d {
		case 0:
			return P{0, 1}
		case 1:
			return P{1, 0}
		case 2:
			return P{0, -1}
		case 3:
			return P{-1, 0}
		default:
			panic("invalid dir")
		}
	}

	state := make([]S, 0)
	for y := range input {
		for x := range input[y] {
			if input[y][x] == 'S' {
				state = append(state, S{p: P{y, x}, d: 0, s: 0})
			}
		}
	}

	pop := func() S {
		s := state[0]
		state[0] = state[len(state)-1]
		state = state[:len(state)-1]
		i := 0

		for {
			if len(state) > (i*2+1) && state[i].s > state[i*2+1].s {
				t := state[i]
				state[i] = state[i*2+1]
				state[i*2+1] = t
				i = i*2 + 1
			} else if len(state) > (i*2+2) && state[i].s > state[i*2+2].s {
				t := state[i]
				state[i] = state[i*2+2]
				state[i*2+2] = t
				i = i*2 + 2
			} else {
				break
			}
		}
		return s
	}

	push := func(s S) {
		i := len(state)
		state = append(state, s)
		for ; i > 0; i = (i - 1) / 2 {
			if state[i].s < state[(i-1)/2].s {
				t := state[i]
				state[i] = state[(i-1)/2]
				state[(i-1)/2] = t
			}
		}
	}

	from := make(map[S][]S)
	visited := make(map[S]int)

	add := func(f, t S) {
		if input[t.p.y][t.p.x] == '#' {
			return
		}
		s, v := visited[S{t.p, t.d, 0}]
		if v && s < t.s {
			return
		}
		visited[S{t.p, t.d, 0}] = t.s

		_, e := from[t]
		if e {
			from[t] = append(from[t], f)
			//fmt.Printf("%v %v %v\n", f, t, from[t])
			//panic("dup")
			return
		}
		//fmt.Printf("from %v => %v\n", f, t)
		a := make([]S, 1)
		a[0] = f
		from[t] = a
		push(t)
	}

	best_score := -1
	end := make([]S, 0)
	for {
		if len(state) == 0 {
			break
		}

		s := pop()
		//fmt.Printf("%v\n", s)

		if best_score != -1 && best_score < s.s {
			break
		}

		if input[s.p.y][s.p.x] == '#' {
			continue
		}
		if input[s.p.y][s.p.x] == 'E' {
			if best_score == -1 {
				if best_score != -1 && best_score != s.s {
					panic("Bad score")
				}
				best_score = s.s
				s.s = best_score
				end = append(end, s)
				continue
			}
		}

		add(s, S{
			p: s.p,
			d: (s.d + 3) % 4,
			s: s.s + 1000,
		})
		add(s, S{
			p: s.p,
			d: (s.d + 1) % 4,
			s: s.s + 1000,
		})
		d := dir(s.d)
		add(s, S{
			p: P{y: s.p.y + d.y, x: s.p.x + d.x},
			d: s.d,
			s: s.s + 1,
		})
	}

	seats := make(map[P]bool)
	for {
		if len(end) == 0 {
			break
		}
		p := end[len(end)-1]
		end = end[:len(end)-1]
		seats[p.p] = true

		//fmt.Printf("from %v: %v\n", p, from[p])

		for _, f := range from[p] {
			end = append(end, f)
		}
	}
	/*
		for y := range input {
			for x := range input[y] {
				if seats[P{y, x}] {
					print("O")
				} else if input[y][x] == '#' {
					print("#")
				} else {
					print(".")
				}
			}
			println()
		}*/

	println("Part 2:", len(seats))
}
