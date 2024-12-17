package main

import (
	"fmt"
	"os"
)

type Machine struct {
	a_x, a_y int
	b_x, b_y int
	p_x, p_y int
}

func main() {
	input := func() []Machine {
		input := make([]Machine, 0)

		file, err := os.Open("day13.in")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for {
			var a_x, a_y int
			var b_x, b_y int
			var p_x, p_y int
			_, err = fmt.Fscanf(file, "Button A: X+%v, Y+%v\nButton B: X+%v, Y+%v\nPrize: X=%v, Y=%v\n\n", &a_x, &a_y, &b_x, &b_y, &p_x, &p_y)
			if err != nil {
				break
			}

			input = append(input, Machine{a_x, a_y, b_x, b_y, p_x, p_y})
		}
		return input
	}()

	part1(input)
	part2(input)
}

func part1(input []Machine) {
	ans := 0

	for _, m := range input {
		low := 0
		for a := range 100 {
			y := m.a_y * a

			if y >= m.p_y {
				break
			}

			b := (m.p_y - y) / m.b_y

			if (m.a_y*a + m.b_y*b) != m.p_y {
				continue
			}
			if (m.a_x*a + m.b_x*b) != m.p_x {
				continue
			}
			val := a*3 + b
			if low == 0 || val < low {
				low = val
			}
		}
		if low != 0 {
			ans += low
		}
	}

	println("Part 1", ans)
}

func part2(input []Machine) {
	ans := 0

	/*
		a_x, a_y = 17, 86
		b_x, b_y = 84, 37
		t_x, t_y = 7870, 6450

		a = (t_y*b_x - t_x*b_y) / (a_y*b_x-a_x*b_y)
		b = (t_x - a*a_x) / b_x

	*/
	for _, m := range input {
		a_x, a_y := m.a_x, m.a_y
		b_x, b_y := m.b_x, m.b_y
		p_x, p_y := m.p_x+10000000000000, m.p_y+10000000000000
		//p_x, p_y := m.p_x, m.p_y

		l := p_y*b_x - p_x*b_y
		f := a_y*b_x - a_x*b_y

		if (l % f) != 0 {
			continue
		}

		a := l / f
		b := (p_x - a*a_x) / b_x
		if (p_x-a*a_x)%b_x != 0 {
			continue
		}

		ans += 3*a + b
	}

	println("Part 2", ans)
}
