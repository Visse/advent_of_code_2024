package main

import (
	"fmt"
	"os"
)

func main() {
	input, moves := func() ([][]byte, []byte) {
		file, err := os.ReadFile("day15.in")
		if err != nil {
			panic(err)
		}

		input := make([][]byte, 0)
		s := 0
		for i := range file {
			if file[i] == '\n' {
				if s == i {
					break
				}
				input = append(input, file[s:i])
				s = i + 1
			}
		}

		return input, file[s:]
	}()

	part1(input, moves)
	part2(input, moves)
}

func part1(input [][]byte, moves []byte) {
	tiles := make([][]byte, len(input))

	p_y, p_x := 0, 0
	for y, r := range input {
		tiles[y] = make([]byte, len(r))
		for x, c := range r {
			tiles[y][x] = c
			if c == '@' {
				p_y, p_x = y, x
			}
		}
	}

	for _, m := range moves {
		d_y, d_x := 0, 0
		switch m {
		case '<':
			d_x = -1
			break
		case '>':
			d_x = 1
			break
		case '^':
			d_y = -1
			break
		case 'v':
			d_y = 1
			break
		default:
			continue
		}

		d := 1
		for tiles[p_y+d*d_y][p_x+d*d_x] == 'O' {
			d += 1
		}
		if tiles[p_y+d*d_y][p_x+d*d_x] == '.' {
			tiles[p_y][p_x] = '.'
			if tiles[p_y+d_y][p_x+d_x] == 'O' {
				tiles[p_y+d*d_y][p_x+d*d_x] = 'O'
			}
			tiles[p_y+d_y][p_x+d_x] = '@'
			p_y += d_y
			p_x += d_x
		}
		/*
			fmt.Printf("%c (%v) [%v %v] %c\n", m, d, d_y, d_x, tiles[p_y+d*d_y][p_x+d*d_x])
			for _, r := range tiles {
				for _, c := range r {
					fmt.Printf("%c", c)
				}
			}
			println()
			println()
		*/
	}

	ans := 0
	for y := range tiles {
		for x := range tiles[y] {
			if tiles[y][x] == 'O' {
				ans += y*100 + x
			}
		}
	}
	println("Part 1:", ans)
}

func part2(input [][]byte, moves []byte) {
	tiles := make([][]byte, len(input))

	p_y, p_x := 0, 0
	for y, r := range input {
		tiles[y] = make([]byte, 2*len(r))
		for x, c := range r {
			switch c {
			case '@':
				p_y, p_x = y, x*2
				tiles[y][x*2+0] = '@'
				tiles[y][x*2+1] = '.'
				break
			case 'O':
				tiles[y][x*2+0] = '['
				tiles[y][x*2+1] = ']'
			default:
				tiles[y][x*2+0] = c
				tiles[y][x*2+1] = c
				break
			}
		}
	}

	for _, m := range moves {
		d_y, d_x := 0, 0
		switch m {
		case '<':
			d_x = -1
			break
		case '>':
			d_x = 1
			break
		case '^':
			d_y = -1
			break
		case 'v':
			d_y = 1
			break
		default:
			continue
		}
		var can_move func(y, x int) bool
		type P struct {
			y, x int
		}
		moved_boxes := make(map[P]bool)
		can_move = func(y, x int) bool {
			if tiles[y][x] == '.' {
				return true
			}
			if tiles[y][x] == '#' {
				return false
			}
			if d_x == -1 && tiles[y][x] == ']' { // horizontal
				moved_boxes[P{y, x - 1}] = true
				return can_move(y, x-2)
			} else if d_x == 1 && tiles[y][x] == '[' {
				moved_boxes[P{y, x}] = true
				return can_move(y, x+2)
			} else if d_y == -1 && (tiles[y][x] == '[' || tiles[y][x] == ']') {
				if tiles[y][x] == ']' {
					x -= 1
				}
				moved_boxes[P{y, x}] = true
				return can_move(y-1, x) && can_move(y-1, x+1)
			} else if d_y == 1 && (tiles[y][x] == '[' || tiles[y][x] == ']') {
				if tiles[y][x] == ']' {
					x -= 1
				}
				moved_boxes[P{y, x}] = true
				return can_move(y+1, x) && can_move(y+1, x+1)
			} else {
				panic("Unreachable")
			}
		}

		if can_move(p_y+d_y, p_x+d_x) {
			for p := range moved_boxes {
				if tiles[p.y][p.x] != '[' || tiles[p.y][p.x+1] != ']' {
					fmt.Printf("%c %c", tiles[p.y][p.x], tiles[p.y][p.x+1])
					panic("bad tile")
				}
				tiles[p.y][p.x] = '.'
				tiles[p.y][p.x+1] = '.'
			}
			for p := range moved_boxes {
				if tiles[p.y+d_y][p.x+d_x] != '.' || tiles[p.y+d_y][p.x+d_x+1] != '.' {
					panic("Occupied!")
				}
				tiles[p.y+d_y][p.x+d_x] = '['
				tiles[p.y+d_y][p.x+d_x+1] = ']'
			}
			tiles[p_y][p_x] = '.'
			tiles[p_y+d_y][p_x+d_x] = '@'
			p_y += d_y
			p_x += d_x
		}

	}
	// for _, r := range tiles {
	// 	for _, c := range r {
	// 		fmt.Printf("%c", c)
	// 	}
	// 	println()
	// }
	// println()
	// println()
	ans := 0
	for y := range tiles {
		for x := range tiles[y] {
			if tiles[y][x] == '[' {
				ans += y*100 + x
			}
		}
	}
	println("Part 2:", ans)

}
