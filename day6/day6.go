package main

import (
	"os"
)

func main() {
	raw_input, err := os.ReadFile("day6.in")
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

func part1(input [][]byte) {
	pos_x, pos_y := 0, 0
	visited := make([][]bool, len(input))
	for y, row := range input {
		visited[y] = make([]bool, len(row))

		for x, c := range row {
			if c == '^' {
				pos_x, pos_y = x, y
			}
		}
	}

	dir_x, dir_y := 0, -1

	for {
		visited[pos_y][pos_x] = true
		pos_x += dir_x
		pos_y += dir_y

		if pos_y < 0 || pos_y >= len(input) {
			break
		}
		row := input[pos_y]
		if pos_x < 0 || pos_x >= len(row) {
			break
		}
		c := row[pos_x]
		if c == '#' {
			pos_x -= dir_x
			pos_y -= dir_y

			tx, ty := dir_x, dir_y
			dir_x = -ty
			dir_y = tx
		}
	}
	// for y, row := range input {
	// 	for x, c := range row {
	// 		if visited[y][x] {
	// 			print("X")
	// 		} else {
	// 			fmt.Printf("%c", c)
	// 		}
	// 	}
	// 	println()
	// }

	ans := 0
	for _, row := range visited {
		for _, c := range row {
			if c {
				ans += 1
			}
		}
	}
	println("Part 1:", ans)
}

type State struct {
	pos_x int
	pos_y int
	dir_x int
	dir_y int
	dir   int
}

func step(state State, input [][]byte) (State, bool) {
	pos_y, pos_x := state.pos_y+state.dir_y, state.pos_x+state.dir_x
	if pos_y < 0 || pos_y >= len(input) || pos_x < 0 || pos_x >= len(input[pos_y]) {
		return state, true
	}
	if input[pos_y][pos_x] == '#' {
		state.dir_y, state.dir_x = state.dir_x, -state.dir_y
		state.dir = (state.dir + 1) % 4
		return state, false
	}
	state.pos_y, state.pos_x = pos_y, pos_x
	return state, false
}

func is_looping(start_y int, start_x int, h int, w int, input [][]byte, visited [][][4]bool) bool {
	for y := range h {
		for x := range w {
			for d := range 4 {
				visited[y][x][d] = false
			}
		}
	}

	state := State{
		pos_x: start_x,
		pos_y: start_y,
		dir_x: 0,
		dir_y: -1,
		dir:   0,
	}
	for {
		next_state, end := step(state, input)
		if end {
			return false
		}
		state = next_state
		if visited[state.pos_y][state.pos_x][state.dir] {
			return true
		}
		visited[state.pos_y][state.pos_x][state.dir] = true
	}
}

func part2(input [][]byte) {
	h, w := len(input), len(input[0])
	start_y, start_x := 0, 0
	visited := make([][][4]bool, len(input))
	walled := make([][]bool, len(input))
	for y := range h {
		visited[y] = make([][4]bool, w)
		walled[y] = make([]bool, w)

		for x := range w {
			if input[y][x] == '^' {
				start_y, start_x = y, x
			}
		}
	}

	state := State{
		pos_x: start_x,
		pos_y: start_y,
		dir_x: 0,
		dir_y: -1,
		dir:   0,
	}
	p_dir := 0
	ans := 0
	for {
		new_state, end := step(state, input)
		if end {
			break
		}
		//fmt.Printf("%v %v\n", state, new_state)
		if p_dir != new_state.dir { // Did we turn?
			p_dir = new_state.dir
			state = new_state
			continue
		}

		if walled[new_state.pos_y][new_state.pos_x] {
			state = new_state
			continue
		}

		// Try to put a wall
		if input[new_state.pos_y][new_state.pos_x] == '.' {
			//println("Checking", new_state.pos_y, new_state.pos_x)
			walled[new_state.pos_y][new_state.pos_x] = true
			input[new_state.pos_y][new_state.pos_x] = '#'
			if is_looping(start_y, start_x, h, w, input, visited) {
				ans += 1
			}
			input[new_state.pos_y][new_state.pos_x] = '.'
		}
		state = new_state

	}
	println("Part 2:", ans)
}
